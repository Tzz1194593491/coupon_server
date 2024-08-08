package main

import (
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/dal"
	couponmeta "github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta/couponmetaservice"
	"github.com/Tzz1194593491/coupon_server/pkg/constants"
	"github.com/Tzz1194593491/coupon_server/pkg/middleware/middleware"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func init() {
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	ip, err := constants.GetOutBoundIP()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", ip+":8889")
	if err != nil {
		panic(err)
	}

	svr := couponmeta.NewServer(new(CouponMetaServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.CouponMetaServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                                   // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr), // address
		server.WithRegistry(r),       // registry
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}