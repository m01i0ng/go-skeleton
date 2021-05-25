// +build wireinject

package di

import (
  "github.com/google/wire"
  "github.com/m01i0ng/go-skeleton/app"
  "github.com/m01i0ng/go-skeleton/internal/config"
  "github.com/m01i0ng/go-skeleton/internal/db"
)

func InitApp() (*app.App, error) {
  panic(wire.Build(config.Provider, db.MysqlProvider, db.RedisProvider, db.MinioProvider, db.MqttProvider, app.NewApp))
}
