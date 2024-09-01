package rpc

import (
	"context"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta/couponmetaservice"
	"github.com/Tzz1194593491/coupon_server/pkg/constants"
	"github.com/Tzz1194593491/coupon_server/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
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
	policy := retry.NewFailurePolicy()
	c, err := couponmetaservice.NewClient(
		constants.CouponMetaServiceName,
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
	couponMetaClient = c
}

func GetCouponValidMetaList(ctx context.Context, req *coupon_meta.GetCouponValidMetaInfoReq) (resp *coupon_meta.GetCouponValidMetaInfoResp, err error) {
	klog.Info(req)
	// 获取券模板
	mataMapResp, err := couponMetaClient.GetCouponValidMetaInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return mataMapResp, nil
}

func TryReduceCouponStock(ctx context.Context, req *coupon_meta.TryReduceCouponStockReq) bool {
	klog.Info(req)
	// 尝试扣减库存
	resp, err := couponMetaClient.TryReduceCouponStock(ctx, req)
	if err != nil {
		return false
	}
	return resp.IsSuccess
}
