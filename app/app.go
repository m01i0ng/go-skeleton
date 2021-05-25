package app

import (
  "database/sql"
  "time"

  mqtt "github.com/eclipse/paho.mqtt.golang"
  "github.com/go-redis/redis/v8"
  "github.com/minio/minio-go/v7"
)

type App struct {
  Redis *redis.Client
  Mysql *sql.DB
  Minio *minio.Client
  Mqtt  mqtt.Client
}

func NewApp(redis *redis.Client, mysql *sql.DB, minio *minio.Client, mqtt mqtt.Client) *App {
  app := &App{
    Redis: redis,
    Mysql: mysql,
    Minio: minio,
    Mqtt:  mqtt,
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
