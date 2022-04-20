package config

import (
	"github.com/go-redis/redis/v8"
)

const RedisAddr = "redis_timer_presets:6379"

func SetupRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		DB:   0,
	})
}
