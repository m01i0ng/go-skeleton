package db

import (
  "database/sql"
  "fmt"

  _ "github.com/go-sql-driver/mysql"
  "github.com/google/wire"
  "github.com/m01i0ng/go-skeleton/internal/config"
)

var MysqlProvider = wire.NewSet(NewMysql)

func NewMysql(c *config.Config) (*sql.DB, error) {
  mc := c.Mysql
  db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mc.Username, mc.Password, mc.Host, mc.Port, mc.Db))
  if err != nil {
    return nil, err
  }

  err = db.Ping()
  if err != nil {
    return nil, err
  }

  return db, nil
}
