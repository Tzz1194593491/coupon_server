package redis

import (
	"github.com/Tzz1194593491/coupon_server/pkg/constants"
	"github.com/redis/go-redis/v9"
	"strings"
)

var RDB *redis.ClusterClient

func Init() {
	RDB = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: constants.RedisCluster,
		NewClient: func(opt *redis.Options) *redis.Client {
			opt.Password = constants.RedisPassword
			opt.ClientName = "redis-" + strings.Split(opt.Addr, ":")[1]
			return redis.NewClient(opt)
		},
	})
}
