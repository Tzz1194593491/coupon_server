package rpc

import (
	"context"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/business_code"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta/couponmetaservice"
	"github.com/Tzz1194593491/coupon_server/pkg/constants"
	"github.com/Tzz1194593491/coupon_server/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var couponMetaClient couponmetaservice.Client

func initCouponMetaRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := couponmetaservice.NewClient(
		constants.CouponMetaServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	couponMetaClient = c
}

func AddCouponMeta(ctx context.Context, req *coupon_meta.AddCouponMetaReq) business_code.BusinessCode {
	resp, err := couponMetaClient.AddCouponMeta(ctx, req)
	if err != nil {
		return business_code.BusinessCode_ADD_FAIL
	}
	if resp.BaseResp.IsError {
		return resp.BaseResp.Code
	}
	return business_code.BusinessCode_SUCCESS
}

func DeleteCouponMeta(ctx context.Context, req *coupon_meta.DeleteCouponMetaReq) business_code.BusinessCode {
	resp, err := couponMetaClient.DeleteCouponMeta(ctx, req)
	if err != nil {
		return business_code.BusinessCode_DELETE_FAIL
	}
	if resp.BaseResp.IsError {
		return resp.BaseResp.Code
	}
	return business_code.BusinessCode_SUCCESS
}

func UpdateCouponMeta(ctx context.Context, req *coupon_meta.UpdateCouponMetaReq) business_code.BusinessCode {
	resp, err := couponMetaClient.UpdateCouponMeta(ctx, req)
	if err != nil {
		return business_code.BusinessCode_UPDATE_FAIL
	}
	if resp.BaseResp.IsError {
		return resp.BaseResp.Code
	}
	return business_code.BusinessCode_SUCCESS
}

func GetCouponMeta(ctx context.Context, req *coupon_meta.GetCouponMetaReq) (business_code.BusinessCode, interface{}) {
	resp, err := couponMetaClient.GetCouponMeta(ctx, req)
	if err != nil {
		return business_code.BusinessCode_GET_FAIL, nil
	}
	if resp.BaseResp.IsError {
		return resp.BaseResp.Code, nil
	}
	return business_code.BusinessCode_SUCCESS, resp
}
