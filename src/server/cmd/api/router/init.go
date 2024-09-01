package router

import (
	"github.com/Tzz1194593491/coupon_server/cmd/api/rpc"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init(r *server.Hertz) {
	// 初始化rpc
	rpc.Init()
	// 加载路由
	couponMetaRouterInit(r)
	couponRecordRouterInit(r)
}
