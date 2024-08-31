package dal

import (
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/dal/db"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/dal/redis"
)

// Init init dal
func Init() {
	db.Init()    // mysql init
	redis.Init() // redis init
}
