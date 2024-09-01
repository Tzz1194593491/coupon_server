package main

import (
	"context"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/service"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/business_code"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"github.com/Tzz1194593491/coupon_server/pkg/pack"
	"github.com/cloudwego/kitex/pkg/klog"
)

// CouponMetaServiceImpl implements the last service interface defined in the IDL.
type CouponMetaServiceImpl struct{}

// GetCouponMeta implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) GetCouponMeta(ctx context.Context, req *coupon_meta.GetCouponMetaReq) (resp *coupon_meta.GetCouponMetaResp, err error) {
	resp = new(coupon_meta.GetCouponMetaResp)
	klog.Info(req)
	err = req.IsValid()
	if err != nil {
		return nil, err
	}
	res, err := service.NewManageCouponMeta(ctx).GetManageCouponMetaByPage(req)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.Fail(business_code.BusinessCode_GET_FAIL)
		return nil, err
	}
	resp.BaseResp = pack.Success(business_code.BusinessCode_SUCCESS)
	resp.CouponMeta = res
	resp.BaseInfo = req.BaseInfo
	return
}

// AddCouponMeta implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) AddCouponMeta(ctx context.Context, req *coupon_meta.AddCouponMetaReq) (resp *coupon_meta.AddCouponMetaResp, err error) {
	resp = new(coupon_meta.AddCouponMetaResp)
	klog.Info(req)
	err = req.IsValid()
	if err != nil {
		return nil, err
	}
	err = service.NewManageCouponMeta(ctx).AddManageCouponMeta(req)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.Fail(business_code.BusinessCode_ADD_FAIL)
		return resp, err
	}
	klog.Info(resp)
	resp.BaseResp = pack.Success(business_code.BusinessCode_SUCCESS)
	return resp, nil
}

// DeleteCouponMeta implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) DeleteCouponMeta(ctx context.Context, req *coupon_meta.DeleteCouponMetaReq) (resp *coupon_meta.DeleteCouponMetaResp, err error) {
	resp = new(coupon_meta.DeleteCouponMetaResp)
	klog.Info(req)
	err = req.IsValid()
	if err != nil {
		return nil, err
	}
	err = service.NewManageCouponMeta(ctx).DeleteManageCouponMeta(req)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.Fail(business_code.BusinessCode_DELETE_FAIL)
		return resp, err
	}
	resp.BaseResp = pack.Success(business_code.BusinessCode_SUCCESS)
	return resp, nil
}

// UpdateCouponMeta implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) UpdateCouponMeta(ctx context.Context, req *coupon_meta.UpdateCouponMetaReq) (resp *coupon_meta.UpdateCouponMetaResp, err error) {
	resp = new(coupon_meta.UpdateCouponMetaResp)
	klog.Info(req)
	err = req.IsValid()
	if err != nil {
		return nil, err
	}
	err = service.NewManageCouponMeta(ctx).UpdateManageCouponMeta(req)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.Fail(business_code.BusinessCode_UPDATE_FAIL)
		return resp, err
	}
	resp.BaseResp = pack.Success(business_code.BusinessCode_SUCCESS)
	return resp, nil
}

// GetCouponValidMetaInfo implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) GetCouponValidMetaInfo(ctx context.Context, req *coupon_meta.GetCouponValidMetaInfoReq) (resp *coupon_meta.GetCouponValidMetaInfoResp, err error) {
	resp = new(coupon_meta.GetCouponValidMetaInfoResp)
	klog.Info(req)
	info, err := service.NewManageCouponMeta(ctx).GetCouponValidMetaInfo(req)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.Fail(business_code.BusinessCode_UPDATE_FAIL)
		return resp, err
	}
	resp.CouponMetaInfo = info
	resp.BaseResp = pack.Success(business_code.BusinessCode_SUCCESS)
	return resp, nil
}

// TryReduceCouponStock implements the CouponMetaServiceImpl interface.
func (s *CouponMetaServiceImpl) TryReduceCouponStock(ctx context.Context, req *coupon_meta.TryReduceCouponStockReq) (resp *coupon_meta.TryReduceCouponStockResp, err error) {
	resp = new(coupon_meta.TryReduceCouponStockResp)
	success := service.NewManageCouponMeta(ctx).TryReduceCouponStock(req)
	resp.IsSuccess = success
	return resp, nil
}
