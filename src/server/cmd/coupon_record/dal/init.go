package dal

import (
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_record/dal/db"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_record/dal/redis"
)

func Init() {
	db.Init()
	redis.Init()
}
