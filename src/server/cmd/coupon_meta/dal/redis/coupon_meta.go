package redis

import (
	"context"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/dal/db"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"github.com/Tzz1194593491/coupon_server/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type CouponMeta struct {
	CouponMetaNo    *int64
	couponMetaType  coupon_meta.CouponMetaType
	CouponMetaStock int32
	ValidStartTime  time.Time
	ValidEndTime    time.Time
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

// ShardingToRedis 将库存分片
func (c *CouponMeta) ShardingToRedis(ctx context.Context) (err error) {
	// 模板基础信息
	couponMetaNo := c.CouponMetaNo
	stock := c.CouponMetaStock
	// 获取redis主节点数量
	masterCount := GetMasterCount()
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
		// 状态机需要做定时kv清理
		err = client.HSet(ctx, oneOfKey,
			&CouponMeta{
				CouponMetaNo:    couponMetaNo,
				CouponMetaStock: int32(oneOfValue),
				ValidStartTime:  c.ValidStartTime,
				ValidEndTime:    c.ValidEndTime,
			}, redis.KeepTTL).Err()
		if err != nil {
			klog.Error(err)
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
	// 构造库存键
	couponMetaKey := utils.GetRedisKey(strconv.FormatInt(*couponMetaNo, 10))
	cmd := RDB.Set(ctx, couponMetaKey, stock, redis.KeepTTL)
	klog.Info(cmd)
	return nil
}

// GetCouponMetaList 根据CouponMeta获取分片
func (c *CouponMeta) GetCouponMetaList(ctx context.Context) map[string]*CouponMeta {
	var list map[string]*CouponMeta
	couponMetaKey := utils.GetRedisKey(strconv.FormatInt(*c.CouponMetaNo, 10))
	iter := RDB.Scan(ctx, 0, couponMetaKey+":"+"*", 0).Iterator()
	for iter.Next(ctx) {
		data := &CouponMeta{}
		key := iter.Val()
		if err := RDB.HGetAll(ctx, key).Scan(&data); err != nil {
			klog.Error(err)
			continue
		}
		list[key] = data
	}
	return list
}
