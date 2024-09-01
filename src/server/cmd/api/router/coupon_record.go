package router

import (
	"github.com/Tzz1194593491/coupon_server/cmd/api/handlers"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func couponRecordRouterInit(r *server.Hertz) {
	v1 := r.Group("/v1")
	couponMeta := v1.Group("/CouponRecord")
	{
		couponMeta.POST("", handlers.SendCoupon)
	}
}
