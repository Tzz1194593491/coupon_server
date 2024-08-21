package pack

import (
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/dal/db"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"github.com/Tzz1194593491/coupon_server/pkg/utils"
)

func CouponMeta(c *db.CouponMeta) *coupon_meta.CouponMeta {
	if c == nil {
		return nil
	}
	return &coupon_meta.CouponMeta{
		CouponMetaNo:   *c.CouponMetaNo,
		Type:           *c.CouponMetaType,
		ValidStartTime: utils.TimeToString(c.ValidStartTime),
		ValidEndTime:   c.ValidEndTime.String(),
		Status:         *c.CouponMetaStatus,
		Stock:          c.CouponMetaStock,
		CreateTime:     utils.TimeToString(c.CreatedAt),
		UpdateTime:     utils.TimeToString(c.UpdatedAt),
		DeleteTime:     utils.TimeToString(c.DeletedAt.Time),
	}
}

func CouponMetas(c []*db.CouponMeta) []*coupon_meta.CouponMeta {
	couponMetas := make([]*coupon_meta.CouponMeta, 0)
	for _, c := range c {
		couponMetas = append(couponMetas, CouponMeta(c))
	}
	return couponMetas
}
