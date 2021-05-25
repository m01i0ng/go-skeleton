package db

import (
  "database/sql"
  "fmt"

  _ "github.com/go-sql-driver/mysql"
  "github.com/google/wire"
  "github.com/kataras/golog"
  "github.com/m01i0ng/go-skeleton/internal/config"
)

var MysqlProvider = wire.NewSet(NewMysql)

func NewMysql(c *config.Config) (*sql.DB, error) {
  m := c.Mysql
  db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.Username, m.Password, m.Host, m.Port, m.Db))
  if err != nil {
    golog.Error(err)
    return nil, err
  }

  err = db.Ping()
  if err != nil {
    golog.Error(err)
    return nil, err
  }

  golog.Infof("Mysql connected to: %s", fmt.Sprintf("%s:%d/%s", m.Host, m.Port, m.Db))

  return db, nil
}
