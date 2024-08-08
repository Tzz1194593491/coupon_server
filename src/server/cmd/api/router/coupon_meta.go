package router

import (
	"github.com/Tzz1194593491/coupon_server/cmd/api/handlers"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func couponMetaRouterInit(r *server.Hertz) {
	r.POST("/CouponMeta/Add", handlers.AddCouponMeta)
}
