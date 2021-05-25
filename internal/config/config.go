package config

import (
  "io/ioutil"
  "os"

  "github.com/goccy/go-yaml"
  "github.com/google/wire"
)

var Provider = wire.NewSet(New)

type Config struct {
  Mysql struct {
    Host     string `yaml:"host"`
    Port     int    `yaml:"port"`
    Db       string `yaml:"db"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
  }
  Redis struct {
    Host     string `yaml:"host"`
    Port     int    `yaml:"port"`
    Db       int    `yaml:"db"`
    Password string `yaml:"password"`
  }
  Minio struct {
    Endpoint  string `yaml:"endpoint"`
    Bucket    string `yaml:"bucket"`
    Tls       bool   `yaml:"tls"`
    AccessKey string `yaml:"access_key"`
    SecretKey string `yaml:"secret_key"`
  }
  Mqtt struct {
    Broker   string `yaml:"broker"`
    User     string `yaml:"user"`
    Password string `yaml:"password"`
  }
}

func New() (*Config, error) {
  file, err := os.Open("config/config.yml")
  if err != nil {
    return nil, err
  }
  defer file.Close()

  content, err := ioutil.ReadAll(file)
  if err != nil {
    return nil, err
  }

  var c Config
  err = yaml.Unmarshal(content, &c)
  if err != nil {
    return nil, err
  }

  return &c, nil
}
