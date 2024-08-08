package main

import (
	"context"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/service"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/business_code"
	coupon_meta "github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"github.com/Tzz1194593491/coupon_server/pkg/pack"
)

// CouponMetaServiceImpl implements the last service interface defined in the IDL.
type CouponMetaServiceImpl struct{}

// GetCouponMeta implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) GetCouponMeta(ctx context.Context, req *coupon_meta.GetCouponMetaReq) (resp *coupon_meta.GetCouponMetaResp, err error) {
	resp = new(coupon_meta.GetCouponMetaResp)
	// todo
	_, err = service.NewManageCouponMeta(ctx).GetManageCouponMetaByPage(req)
	if err != nil {
		resp.BaseResp = pack.Fail(business_code.BusinessCode_GET_FAIL)
		return nil, err
	}
	resp.BaseResp = pack.Success(business_code.BusinessCode_SUCCESS)
	//resp.CouponMeta =
	return
}

// AddCouponMeta implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) AddCouponMeta(ctx context.Context, req *coupon_meta.AddCouponMetaReq) (resp *coupon_meta.AddCouponMetaResp, err error) {
	resp = new(coupon_meta.AddCouponMetaResp)

	err = service.NewManageCouponMeta(ctx).AddManageCouponMeta(req)
	if err != nil {
		resp.BaseResp = pack.Fail(business_code.BusinessCode_ADD_FAIL)
		return resp, nil
	}
	resp.BaseResp = pack.Success(business_code.BusinessCode_SUCCESS)
	return resp, nil
}

// DeleteCouponMeta implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) DeleteCouponMeta(ctx context.Context, req *coupon_meta.DeleteCouponMetaReq) (resp *coupon_meta.DeleteCouponMetaResp, err error) {
	resp = new(coupon_meta.DeleteCouponMetaResp)

	err = service.NewManageCouponMeta(ctx).DeleteManageCouponMeta(req)
	if err != nil {
		resp.BaseResp = pack.Fail(business_code.BusinessCode_DELETE_FAIL)
		return resp, nil
	}
	resp.BaseResp = pack.Success(business_code.BusinessCode_SUCCESS)
	return resp, nil
}

// UpdateCouponMeta implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) UpdateCouponMeta(ctx context.Context, req *coupon_meta.UpdateCouponMetaReq) (resp *coupon_meta.UpdateCouponMetaResp, err error) {
	resp = new(coupon_meta.UpdateCouponMetaResp)

	err = service.NewManageCouponMeta(ctx).UpdateManageCouponMeta(req)
	if err != nil {
		resp.BaseResp = pack.Fail(business_code.BusinessCode_UPDATE_FAIL)
		return resp, nil
	}
	resp.BaseResp = pack.Success(business_code.BusinessCode_SUCCESS)
	return resp, nil
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
