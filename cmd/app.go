package main

import (
  "database/sql"
  "time"

  "github.com/go-redis/redis/v8"
)

type App struct {
  redis *redis.Client
  mysql *sql.DB
}

func NewApp(redis *redis.Client, mysql *sql.DB) *App {
  app := &App{
    redis: redis,
    mysql: mysql,
  }

  app.firstHeartBeat()

  go func() {
    timer := time.NewTicker(time.Second * 2)
    for {
      app.heartbeat()
      <-timer.C
    }
  }()

  return app
}

func (a *App) firstHeartBeat() error {
  return nil
}

func (a *App) heartbeat() {
}
