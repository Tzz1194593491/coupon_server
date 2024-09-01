package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	Init()
}

func TestIncrBy(t *testing.T) {
	key := "1"
	value := 2
	// 初始化kv{"1": 1}
	RDB.Set(context.Background(), key, value, redis.KeepTTL)
	keys := []string{key}
	values := []interface{}{-1 * value}
	// 第一次扣减成功，返回1
	num, _ := StockReduce.Run(context.Background(), RDB, keys, values).Int()
	assert.Equal(t, 1, num)
	// 第二次扣减失败，返回0
	num, _ = StockReduce.Run(context.Background(), RDB, keys, values).Int()
	assert.Equal(t, 0, num)
	// 不存在的kv
	num, _ = StockReduce.Run(context.Background(), RDB, []string{"233"}, values).Int()
	assert.Equal(t, 0, num)
}
