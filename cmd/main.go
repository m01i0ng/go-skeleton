package main

import "github.com/m01i0ng/go-skeleton/di"

func main() {
  app, err := di.InitApp()
  if err != nil {
    return
  }
  app.Redis.String()

  select {}
}
