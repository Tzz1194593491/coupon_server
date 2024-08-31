package redis

import (
	"context"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/dal/db"
	"github.com/Tzz1194593491/coupon_server/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type CouponMeta struct {
	CouponMetaNo    *int64
	CouponMetaStock int32
}

func WithCouponMeta(couponMeta *db.CouponMeta) *CouponMeta {
	return &CouponMeta{
		CouponMetaNo:    couponMeta.CouponMetaNo,
		CouponMetaStock: couponMeta.CouponMetaStock,
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
		set := client.Set(ctx, oneOfKey, oneOfValue, redis.KeepTTL)
		klog.Info(set)
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
