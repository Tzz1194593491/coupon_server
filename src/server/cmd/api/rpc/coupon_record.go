package rpc

import (
	"context"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/business_code"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_record"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_record/couponrecordservice"
	"github.com/Tzz1194593491/coupon_server/pkg/constants"
	"github.com/Tzz1194593491/coupon_server/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var couponRecordClient couponrecordservice.Client

func initCouponRecordRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	policy := retry.NewFailurePolicy()
	c, err := couponrecordservice.NewClient(
		constants.CouponRecordServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                    // mux
		client.WithRPCTimeout(3*time.Second),           // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond), // conn timeout
		client.WithFailureRetry(policy),                // retry
		client.WithResolver(r),                         // resolver
	)
	if err != nil {
		panic(err)
	}
	couponRecordClient = c
}

func SendCoupon(ctx context.Context, req *coupon_record.CouponRecordSendReq) business_code.BusinessCode {
	_, err := couponRecordClient.SendCoupon(ctx, req)
	if err != nil {
		return business_code.BusinessCode_SEND_COUPON_FAIL
	}
	return business_code.BusinessCode_SUCCESS
}
