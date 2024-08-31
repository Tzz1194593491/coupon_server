package main

import (
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/dal"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/utils"
	couponmeta "github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta/couponmetaservice"
	"github.com/Tzz1194593491/coupon_server/pkg/constants"
	middleware2 "github.com/Tzz1194593491/coupon_server/pkg/middleware"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func init() {
	dal.Init()
	utils.Init()
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
	addr, err := net.ResolveTCPAddr("tcp", ip+":"+constants.CouponMetaServicePort)
	if err != nil {
		panic(err)
	}

	svr := couponmeta.NewServer(new(CouponMetaServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.CouponMetaServiceName}), // server name
		server.WithMiddleware(middleware2.CommonMiddleware),                                                  // middleware
		server.WithMiddleware(middleware2.ServerMiddleware),
		server.WithServiceAddr(addr), // address
		server.WithRegistry(r),       // registry
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
