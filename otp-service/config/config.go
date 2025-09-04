package config

import (
  "context"
  "fmt"
  "os"

  "github.com/go-redis/redis/v8"
)

var (
  Ctx         = context.Background()
  RedisClient *redis.Client
)

func InitRedis() {
  // read host & port from env, default to localhost:6379 for non-Docker usage
  host := os.Getenv("REDIS_HOST")
  if host == "" {
    host = "localhost"
  }

  port := os.Getenv("REDIS_PORT")
  if port == "" {
    port = "6379"
  }

  addr := fmt.Sprintf("%s:%s", host, port)
  RedisClient = redis.NewClient(&redis.Options{
    Addr: addr,
  })

  if _, err := RedisClient.Ping(Ctx).Result(); err != nil {
    panic(fmt.Sprintf("Failed to connect to Redis at %s: %v", addr, err))
  }
}