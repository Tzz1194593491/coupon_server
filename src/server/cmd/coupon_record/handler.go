package main

import (
	"context"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_record/service"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/business_code"
	coupon_record "github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_record"
	"github.com/Tzz1194593491/coupon_server/pkg/pack"
)

// CouponRecordServiceImpl implements the last service interface defined in the IDL.
type CouponRecordServiceImpl struct{}

// SendCoupon implements the CouponRecordServiceImpl interface.
func (s *CouponRecordServiceImpl) SendCoupon(ctx context.Context, req *coupon_record.CouponRecordSendReq) (resp *coupon_record.CouponRecordSendResp, err error) {
	resp = new(coupon_record.CouponRecordSendResp)
	err = service.NewCouponRecordService(ctx).SendCoupon(req)
	if err != nil {
		return nil, err
	}
	resp.CouponMetaNo = req.GetCouponMetaNo()
	resp.UserId = req.UserId
	resp.BaseResp = pack.Success(business_code.BusinessCode_SUCCESS)
	return resp, nil
}

// GetCouponRecordList implements the CouponRecordServiceImpl interface.
func (s *CouponRecordServiceImpl) GetCouponRecordList(ctx context.Context, req *coupon_record.CouponRecordListReq) (resp *coupon_record.CouponRecordListResp, err error) {
	return
}

// UseCoupon implements the CouponRecordServiceImpl interface.
func (s *CouponRecordServiceImpl) UseCoupon(ctx context.Context, req *coupon_record.CouponRecordUsedReq) (resp *coupon_record.CouponRecordUsedResp, err error) {
	return
}
