package main

import (
	"context"
	coupon_meta "github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"github.com/Tzz1194593491/coupon_server/rpc/coupon_meta/pkg/mysql"
)

// CouponMetaServiceImpl implements the last service interface defined in the IDL.
type CouponMetaServiceImpl struct {
	couponMetaManager *mysql.CouponMetaManager
}

// GetCouponMeta implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) GetCouponMeta(ctx context.Context, req *coupon_meta.GetCouponMetaReq) (resp *coupon_meta.GetCouponMetaResp, err error) {
	// TODO: Your code here...
	return
}

// AddCouponMeta implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) AddCouponMeta(ctx context.Context, req *coupon_meta.AddCouponMetaReq) (resp *coupon_meta.AddCouponMetaResp, err error) {
	// TODO: Your code here...
	return
}

// DeleteCouponMeta implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) DeleteCouponMeta(ctx context.Context, req *coupon_meta.DeleteCouponMetaReq) (resp *coupon_meta.DeleteCouponMetaResp, err error) {
	// TODO: Your code here...
	return
}

// UpdateCouponMeta implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) UpdateCouponMeta(ctx context.Context, req *coupon_meta.UpdateCouponMetaReq) (resp *coupon_meta.UpdateCouponMetaResp, err error) {
	// TODO: Your code here...
	return
}

// GetCouponMetaIsValid implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) GetCouponMetaIsValid(ctx context.Context, res *coupon_meta.GetCouponMetaIsValidResp) (resp *coupon_meta.GetCouponMetaIsValidReq, err error) {
	// TODO: Your code here...
	return
}

// GetCouponMetaStock implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) GetCouponMetaStock(ctx context.Context, res *coupon_meta.GetCouponMetaStockReq) (resp *coupon_meta.GetCouponMetaStockResp, err error) {
	// TODO: Your code here...
	return
}
