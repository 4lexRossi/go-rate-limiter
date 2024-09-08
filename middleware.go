package main

import (
	"net/http"
	"os"
	"strconv"
	"time"
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	ipLimit, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_IP"))
	tokenLimit, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_TOKEN"))
	blockDuration, _ := strconv.Atoi(os.Getenv("BLOCK_DURATION"))

	redisClient := getRedisClient()
	rateLimiter := NewRateLimiter(redisClient, ipLimit, tokenLimit, time.Duration(blockDuration)*time.Second)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		token := r.Header.Get("API_KEY")

		if !rateLimiter.AllowRequest(ip, token) {
			http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
