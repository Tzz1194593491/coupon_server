// Code generated by Kitex v0.10.3. DO NOT EDIT.

package couponmetaservice

import (
	"context"
	coupon_meta "github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetCouponMeta(ctx context.Context, req *coupon_meta.GetCouponMetaReq, callOptions ...callopt.Option) (r *coupon_meta.GetCouponMetaResp, err error)
	AddCouponMeta(ctx context.Context, req *coupon_meta.AddCouponMetaReq, callOptions ...callopt.Option) (r *coupon_meta.AddCouponMetaResp, err error)
	DeleteCouponMeta(ctx context.Context, req *coupon_meta.DeleteCouponMetaReq, callOptions ...callopt.Option) (r *coupon_meta.DeleteCouponMetaResp, err error)
	UpdateCouponMeta(ctx context.Context, req *coupon_meta.UpdateCouponMetaReq, callOptions ...callopt.Option) (r *coupon_meta.UpdateCouponMetaResp, err error)
	GetCouponValidMetaInfo(ctx context.Context, req *coupon_meta.GetCouponValidMetaInfoReq, callOptions ...callopt.Option) (r *coupon_meta.GetCouponValidMetaInfoResp, err error)
	TryReduceCouponStock(ctx context.Context, req *coupon_meta.TryReduceCouponStockReq, callOptions ...callopt.Option) (r *coupon_meta.TryReduceCouponStockResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kCouponMetaServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCouponMetaServiceClient struct {
	*kClient
}

func (p *kCouponMetaServiceClient) GetCouponMeta(ctx context.Context, req *coupon_meta.GetCouponMetaReq, callOptions ...callopt.Option) (r *coupon_meta.GetCouponMetaResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCouponMeta(ctx, req)
}

func (p *kCouponMetaServiceClient) AddCouponMeta(ctx context.Context, req *coupon_meta.AddCouponMetaReq, callOptions ...callopt.Option) (r *coupon_meta.AddCouponMetaResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddCouponMeta(ctx, req)
}

func (p *kCouponMetaServiceClient) DeleteCouponMeta(ctx context.Context, req *coupon_meta.DeleteCouponMetaReq, callOptions ...callopt.Option) (r *coupon_meta.DeleteCouponMetaResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCouponMeta(ctx, req)
}

func (p *kCouponMetaServiceClient) UpdateCouponMeta(ctx context.Context, req *coupon_meta.UpdateCouponMetaReq, callOptions ...callopt.Option) (r *coupon_meta.UpdateCouponMetaResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateCouponMeta(ctx, req)
}

func (p *kCouponMetaServiceClient) GetCouponValidMetaInfo(ctx context.Context, req *coupon_meta.GetCouponValidMetaInfoReq, callOptions ...callopt.Option) (r *coupon_meta.GetCouponValidMetaInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCouponValidMetaInfo(ctx, req)
}

func (p *kCouponMetaServiceClient) TryReduceCouponStock(ctx context.Context, req *coupon_meta.TryReduceCouponStockReq, callOptions ...callopt.Option) (r *coupon_meta.TryReduceCouponStockResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TryReduceCouponStock(ctx, req)
}
