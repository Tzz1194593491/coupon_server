package router

import (
	"github.com/Tzz1194593491/coupon_server/cmd/api/handlers"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func couponMetaRouterInit(r *server.Hertz) {
	v1 := r.Group("/v1")
	couponMeta := v1.Group("/CouponMeta")
	{
		couponMeta.POST("", handlers.AddCouponMeta)
		couponMeta.DELETE("", handlers.DeleteCouponMeta)
		couponMeta.PUT("", handlers.UpdateCouponMeta)
		couponMeta.GET("/ByPage", handlers.GetCouponMetaByPage)
	}
}
