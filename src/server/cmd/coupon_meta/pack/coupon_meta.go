package pack

import (
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/dal/db"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/dal/redis"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"github.com/Tzz1194593491/coupon_server/pkg/utils"
)

func DB2CouponMeta(c *db.CouponMeta) *coupon_meta.CouponMeta {
	if c == nil {
		return nil
	}
	return &coupon_meta.CouponMeta{
		CouponMetaNo:   *c.CouponMetaNo,
		Type:           *c.CouponMetaType,
		ValidStartTime: utils.TimeToString(c.ValidStartTime),
		ValidEndTime:   utils.TimeToString(c.ValidEndTime),
		Status:         *c.CouponMetaStatus,
		Stock:          c.CouponMetaStock,
		CreateTime:     utils.TimeToString(c.CreatedAt),
		UpdateTime:     utils.TimeToString(c.UpdatedAt),
		DeleteTime:     utils.TimeToString(c.DeletedAt.Time),
	}
}

func DB2CouponMetas(c []*db.CouponMeta) []*coupon_meta.CouponMeta {
	couponMetas := make([]*coupon_meta.CouponMeta, 0)
	for _, c := range c {
		couponMetas = append(couponMetas, DB2CouponMeta(c))
	}
	return couponMetas
}

func Redis2CouponMeta(c *redis.CouponMeta) *coupon_meta.CouponMeta {
	if c == nil {
		return nil
	}
	return &coupon_meta.CouponMeta{
		CouponMetaNo:   *c.CouponMetaNo,
		ValidStartTime: utils.TimeToString(c.ValidStartTime),
		ValidEndTime:   utils.TimeToString(c.ValidEndTime),
		Stock:          c.CouponMetaStock,
	}
}
