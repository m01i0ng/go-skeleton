package db

import (
  "context"
  "fmt"

  "github.com/go-redis/redis/v8"
  "github.com/google/wire"
  "github.com/m01i0ng/go-skeleton/internal/config"
)

var RedisProvider = wire.NewSet(NewRedis)

func NewRedis(c *config.Config) (*redis.Client, error) {
  rc := c.Redis
  client := redis.NewClient(&redis.Options{
    Addr:     fmt.Sprintf("%s:%d", rc.Host, rc.Port),
    DB:       rc.Db,
    Password: rc.Password,
  })

  ctx := context.Background()
  if ping := client.Ping(ctx); ping.Err() != nil {
    return nil, ping.Err()
  }

  return client, nil
}
