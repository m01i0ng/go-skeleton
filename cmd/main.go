package main

import (
  "github.com/kataras/golog"
  "github.com/m01i0ng/go-skeleton/di"
)

func main() {
  app, err := di.InitApp()
  if err != nil {
    golog.Fatal(err)
  }

  app.Mqtt.IsConnected()
  select {}
}
