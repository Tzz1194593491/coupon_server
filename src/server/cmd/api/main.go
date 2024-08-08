package main

import (
	"github.com/Tzz1194593491/coupon_server/cmd/api/middleware"
	"github.com/Tzz1194593491/coupon_server/cmd/api/router"
	"github.com/Tzz1194593491/coupon_server/pkg/constants"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init(r *server.Hertz) {
	middleware.Init(r)
	router.Init(r)
}

func main() {
	ip, err := constants.GetOutBoundIP()
	if err != nil {
		panic(err)
	}
	r := server.New(
		server.WithHostPorts(ip+":"+constants.ApiServicePort),
		server.WithHandleMethodNotAllowed(true),
	)
	Init(r)
	r.Spin()
}
