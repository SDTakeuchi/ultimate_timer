package config

import (
	"fmt"
	"os"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

func NewRedis() *redis.Client {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
	reddisAddr := fmt.Sprintf(
		"%s:%s",
		os.Getenv("REDIS_ADDRESS"),
		os.Getenv("REDIS_PORT"),
	)
	return redis.NewClient(&redis.Options{
		Addr: reddisAddr,
		DB:   0,
	})
}
