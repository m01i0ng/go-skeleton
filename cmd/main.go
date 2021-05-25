package main

func main() {
  app, err := InitApp()
  if err != nil {
    return
  }
  app.redis.String()
  executor := NewExecutor(app)
  executor.Exec()
  select {}
}
