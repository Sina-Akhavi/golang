package config

import (
    "github.com/go-redis/redis/v8"
    "context"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis server address
    })

    _, err := RedisClient.Ping(Ctx).Result()
    if err != nil {
        panic("Failed to connect to Redis: " + err.Error())
    }
}