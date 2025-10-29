package commons

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func RedisClientConnect() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-12372.c278.us-east-1-4.ec2.redns.redis-cloud.com:12372",
		Username: "default",
		Password: "F3Y7NKWsnrjNvW1vN8UkZzhRZiiwTCm8",
		DB:       0,
	})

	return rdb

}

func SetKeyValue(ctx context.Context, rdb *redis.Client, key string, value string, ttl time.Duration) *redis.StatusCmd {
	res := rdb.Set(ctx, key, value, ttl)
	return res
}

func GetValueByKey(ctx context.Context, rdb *redis.Client, key string) *redis.StringCmd {
	res := rdb.Get(ctx, key)
	return res
}
