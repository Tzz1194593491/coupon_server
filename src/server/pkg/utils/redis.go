package utils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strings"
)

const (
	PreKey = "CouponMeta:"
)

func GetRedisKey(keys ...string) string {
	key := PreKey
	return key + strings.Join(keys, ":")
}

func ShareKey(key string, masterCount int) *[]string {
	keyList := make([]string, masterCount)
	for i := 0; i < masterCount; i++ {
		keyList[i] = fmt.Sprintf(key+":"+"%d", i)
	}
	return &keyList
}

func GetMasterCount(RDB *redis.ClusterClient) int {
	ctx := context.Background()
	nodes := RDB.ClusterNodes(ctx)
	split := strings.Split(nodes.Val(), "\n")
	master := 0
	for _, value := range split {
		i := strings.Index(value, "master")
		if i != -1 {
			master++
		}
	}
	return master
}
