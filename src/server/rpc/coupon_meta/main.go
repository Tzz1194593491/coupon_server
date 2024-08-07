package main

import (
	coupon_meta "github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta/couponmetaservice"
	"log"
)

func main() {
	svr := coupon_meta.NewServer(new(CouponMetaServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
