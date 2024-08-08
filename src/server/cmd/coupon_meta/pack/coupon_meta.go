package pack

import (
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/dal/db"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
)

func CouponMeta(c *db.CouponMeta) *coupon_meta.CouponMeta {
	if c == nil {
		return nil
	}
	return &coupon_meta.CouponMeta{
		CouponMetaNo:   c.CouponMetaNo,
		Type:           c.CouponMetaType,
		ValidStartTime: c.ValidStartTime.Unix(),
		ValidEndTime:   c.ValidEndTime.Unix(),
		Status:         c.CouponMetaStatus,
		Stock:          c.CouponMetaStock,
		CreateTime:     c.CreateTime.Unix(),
		UpdateTime:     c.UpdateTime.Unix(),
		DeleteTime:     c.Deleted.Time.Unix(),
	}
}

func CouponMetas(c []*db.CouponMeta) []*coupon_meta.CouponMeta {
	couponMetas := make([]*coupon_meta.CouponMeta, 0)
	for _, c := range c {
		couponMetas = append(couponMetas, CouponMeta(c))
	}
	return couponMetas
}
