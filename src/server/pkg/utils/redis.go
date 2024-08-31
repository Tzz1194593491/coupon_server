package utils

import (
	"fmt"
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
