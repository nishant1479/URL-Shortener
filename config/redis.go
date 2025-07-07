package config

import (
    "context"
    "github.com/redis/go-redis/v9"
    "log"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", // or from .env
        Password: "",               // no password set
        DB:       0,
    })

    _, err := RedisClient.Ping(Ctx).Result()
    if err != nil {
        log.Fatalf("Redis connection failed: %v", err)
    }

    log.Println("The connection to Redis has been established")
}
