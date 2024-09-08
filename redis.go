package main

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var redisClient *redis.Client

func initRedis() {
	redisAddr := os.Getenv("REDIS_ADDR")
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
}

func getRedisClient() *redis.Client {
	return redisClient
}
