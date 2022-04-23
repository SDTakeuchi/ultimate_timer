package config

import (
	"github.com/go-redis/redis/v8"
)

// TODO: move to env file
const RedisAddr = "redis_timer_presets:6379"

func NewRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		DB:   0,
	})
}
