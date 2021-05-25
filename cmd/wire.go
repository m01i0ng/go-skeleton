// +build wireinject

package main

import (
  "github.com/google/wire"
  "github.com/m01i0ng/go-skeleton/internal/config"
  "github.com/m01i0ng/go-skeleton/internal/db"
)

func InitApp() (*App, error) {
  panic(wire.Build(config.Provider, db.MysqlProvider, db.RedisProvider, NewApp))
}
