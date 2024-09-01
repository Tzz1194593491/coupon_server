package redis

import (
	"context"
	"encoding/json"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/dal/db"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"github.com/Tzz1194593491/coupon_server/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

const (
	STOCK_REDUCE_SUCCESS = 1
	COUPON_META_INFO     = "info"
)

var StockReduce = redis.NewScript(`
	local key = tostring(KEYS[1])
	local change = tonumber(ARGV[1])
	local value = tonumber(redis.call("GET", key))
	local result = 1
	if not value then
		return 0;
	end
	if value <= 0 then
		return 0;
	end
	redis.call("INCRBY", key, change)
	return 1;
`)

type CouponMeta struct {
	CouponMetaNo    *int64                     `redis:"CouponMetaNo"`
	CouponMetaType  coupon_meta.CouponMetaType `redis:"CouponMetaType"`
	CouponMetaStock int32                      `redis:"CouponMetaStock"`
	ValidStartTime  time.Time                  `redis:"ValidStartTime"`
	ValidEndTime    time.Time                  `redis:"ValidEndTime"`
}

func WithCouponMeta(couponMeta *db.CouponMeta) *CouponMeta {
	return &CouponMeta{
		CouponMetaNo:    couponMeta.CouponMetaNo,
		CouponMetaStock: couponMeta.CouponMetaStock,
		ValidStartTime:  couponMeta.ValidStartTime,
		ValidEndTime:    couponMeta.ValidEndTime,
	}
}

func WithGetCouponMeta(couponMetaNo int64) *CouponMeta {
	return &CouponMeta{
		CouponMetaNo: &couponMetaNo,
	}
}

func WithReduceCouponStock(couponMetaNo int64) *CouponMeta {
	return &CouponMeta{
		CouponMetaNo: &couponMetaNo,
	}
}

// ShardingToRedis 将库存分片
func (c *CouponMeta) ShardingToRedis(ctx context.Context) (err error) {
	// 模板基础信息
	couponMetaNo := c.CouponMetaNo
	stock := c.CouponMetaStock
	// 获取redis主节点数量
	masterCount := utils.GetMasterCount(RDB)
	// 构造库存键
	couponMetaKey := utils.GetRedisKey(strconv.FormatInt(*couponMetaNo, 10))
	couponMetaKeyList := *utils.ShareKey(couponMetaKey, masterCount)
	// 非均匀均分
	equipartition := *utils.Equipartition(int(stock), masterCount)
	// 分发库存
	for i := 0; i < masterCount; i++ {
		oneOfKey := couponMetaKeyList[i]
		oneOfValue := equipartition[i]
		client, err := RDB.MasterForKey(ctx, oneOfKey)
		if err != nil {
			return err
		}
		marshal, err := json.Marshal(oneOfValue)
		if err != nil {
			return err
		}
		// 状态机需要做定时kv清理
		err = client.Set(ctx, oneOfKey, marshal, redis.KeepTTL).Err()
		if err != nil {
			klog.Error(err)
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}

// SingleToRedis 不对库存进行分片处理
func (c *CouponMeta) SingleToRedis(ctx context.Context) error {
	// 模板基础信息
	couponMetaNo := c.CouponMetaNo
	stock := c.CouponMetaStock
	marshal, err := json.Marshal(stock)
	if err != nil {
		return err
	}
	// 构造库存键
	couponMetaKey := utils.GetRedisKey(strconv.FormatInt(*couponMetaNo, 10))
	cmd := RDB.Set(ctx, couponMetaKey, marshal, redis.KeepTTL)
	klog.Info(cmd)
	return nil
}

// GetCouponMetaInfo 获取券模板信息
func (c *CouponMeta) GetCouponMetaInfo(ctx context.Context) (res *CouponMeta, err error) {
	couponMetaKey := utils.GetRedisKey(COUPON_META_INFO, strconv.FormatInt(*c.CouponMetaNo, 10))
	if err := RDB.HGetAll(ctx, couponMetaKey).Scan(res); err != nil {
		return nil, err
	}
	return res, nil
}

// WriteToCouponMetaInfo 写入券模板信息
func (c *CouponMeta) WriteToCouponMetaInfo(ctx context.Context, redis *CouponMeta) error {
	couponMetaKey := utils.GetRedisKey(COUPON_META_INFO, strconv.FormatInt(*c.CouponMetaNo, 10))
	couponMetaNoM, _ := json.Marshal(redis.CouponMetaNo)
	couponMetaTypeM, _ := json.Marshal(redis.CouponMetaType)
	couponMetaStockM, _ := json.Marshal(redis.CouponMetaStock)
	validStartTimeM, _ := json.Marshal(redis.ValidStartTime)
	validEndTimeM, _ := json.Marshal(redis.ValidEndTime)
	pipeline := RDB.Pipeline()
	pipeline.HSet(ctx, couponMetaKey, "CouponMetaNo", couponMetaNoM)
	pipeline.HSet(ctx, couponMetaKey, "CouponMetaType", couponMetaTypeM)
	pipeline.HSet(ctx, couponMetaKey, "CouponMetaStock", couponMetaStockM)
	pipeline.HSet(ctx, couponMetaKey, "ValidStartTime", validStartTimeM)
	pipeline.HSet(ctx, couponMetaKey, "ValidEndTime", validEndTimeM)
	pipeline.Expire(ctx, couponMetaKey, 20*time.Minute)
	_, err := pipeline.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// GetCouponMetaList 根据CouponMeta获取分片
func (c *CouponMeta) GetCouponMetaList(ctx context.Context) *[]string {
	res := make([]string, 0)
	couponMetaKey := utils.GetRedisKey(strconv.FormatInt(*c.CouponMetaNo, 10))
	err := RDB.ForEachMaster(ctx, func(ctx context.Context, rdb *redis.Client) error {
		iter := rdb.Scan(ctx, 0, couponMetaKey+":"+"*", 0).Iterator()
		for iter.Next(ctx) {
			key := iter.Val()
			res = append(res, key)
		}
		return iter.Err()
	})
	if err != nil {
		return nil
	}
	return &res
}

// ReduceStock 扣减库存
func (c *CouponMeta) ReduceStock(ctx context.Context, couponMetaKeys *[]string) bool {
	keys := *couponMetaKeys
	// 打乱数组
	utils.Shuffle(keys)
	// 依次询问扣减
	for _, key := range keys {
		keys := []string{key}
		values := []interface{}{-1}
		num, err := StockReduce.Run(ctx, RDB, keys, values).Int()
		if err != nil {
			return false
		}
		if num == STOCK_REDUCE_SUCCESS {
			return true
		}
	}
	return false
}
