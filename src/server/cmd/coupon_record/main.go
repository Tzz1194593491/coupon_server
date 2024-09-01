package main

import (
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_record/dal"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_record/rpc"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_record/tools"
	couponrecord "github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_record/couponrecordservice"
	"github.com/Tzz1194593491/coupon_server/pkg/constants"
	"github.com/Tzz1194593491/coupon_server/pkg/middleware"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func init() {
	dal.Init()
	rpc.Init()
	tools.Init()
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
	addr, err := net.ResolveTCPAddr("tcp", ip+":"+constants.CouponRecordServicePort)
	if err != nil {
		panic(err)
	}

	svr := couponrecord.NewServer(new(CouponRecordServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.CouponRecordServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                                     // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr), // address
		server.WithRegistry(r),       // registry
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
