package service

import (
	"context"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/dal/db"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"time"
)

type StatusCouponMeta struct {
	ctx context.Context
}

func NewStatusCouponMeta(ctx context.Context) *StatusCouponMeta {
	return &StatusCouponMeta{
		ctx: ctx,
	}
}

func (s *StatusCouponMeta) GetStatusCouponMetaIsValid(req *coupon_meta.GetCouponMetaIsValidReq) (isValid bool) {
	one, err := db.GetCouponMetaById(s.ctx, &db.CouponMeta{
		CouponMetaNo: &req.CouponMetaNo,
	})
	if err != nil {
		return false
	}
	isValid = one.ValidStartTime.Before(time.Now()) && one.ValidEndTime.After(time.Now())
	return isValid
}

func (s *StatusCouponMeta) GetStatusCouponMetaStock(req *coupon_meta.GetCouponMetaStockReq) (stock int32) {
	one, err := db.GetCouponMetaById(s.ctx, &db.CouponMeta{
		CouponMetaNo: &req.CouponMetaNo,
	})
	if err != nil {
		return 0
	}
	return one.CouponMetaStock
}
