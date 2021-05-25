package main

type Executor struct {
  app *App
}

func NewExecutor(app *App) *Executor {
  e := &Executor{app: app}
  return e
}

func (e *Executor) Exec() {

}
