package main

import (
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	rateLimitIPKey    = "rate_limit_ip:"
	rateLimitTokenKey = "rate_limit_token:"
)

type RateLimiter struct {
	redisClient   *redis.Client
	ipLimit       int
	tokenLimit    int
	blockDuration time.Duration
}

func NewRateLimiter(redisClient *redis.Client, ipLimit, tokenLimit int, blockDuration time.Duration) *RateLimiter {
	return &RateLimiter{
		redisClient:   redisClient,
		ipLimit:       ipLimit,
		tokenLimit:    tokenLimit,
		blockDuration: blockDuration,
	}
}

func (rl *RateLimiter) AllowRequest(ip, token string) bool {
	now := time.Now().Unix()

	ipKey := rateLimitIPKey + ip
	ipRequests, _ := rl.redisClient.LRange(ctx, ipKey, 0, -1).Result()
	if len(ipRequests) >= rl.ipLimit {
		return false
	}
	rl.redisClient.RPush(ctx, ipKey, now)
	rl.redisClient.Expire(ctx, ipKey, rl.blockDuration*time.Second)

	tokenKey := rateLimitTokenKey + token
	tokenRequests, _ := rl.redisClient.LRange(ctx, tokenKey, 0, -1).Result()
	if len(tokenRequests) >= rl.tokenLimit {
		return false
	}
	rl.redisClient.RPush(ctx, tokenKey, now)
	rl.redisClient.Expire(ctx, tokenKey, rl.blockDuration*time.Second)

	return true
}
