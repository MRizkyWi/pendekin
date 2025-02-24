package cache

import (
	"context"
	"fmt"
	"log"
	"pendekin/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client

	defaultExpiryTime = 24 * time.Hour
)

func ConnectRedis() {
	redisHost := config.GetEnv("REDIS_HOST", "http://localhost")
	redisPort := config.GetEnv("REDIS_PORT", "6379")

	opt, err := redis.ParseURL(fmt.Sprintf("redis://%s:%s", redisHost, redisPort))
	if err != nil {
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}

	redisClient = redis.NewClient(opt)

	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Redis connected!")
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func GetDefaultExpiryTime() time.Duration {
	return defaultExpiryTime
}