package main

import (
  "github.com/m01i0ng/go-skeleton/app"
)

type Executor struct {
  app *app.App
}

func NewExecutor(app *app.App) (*Executor, error) {
  e := &Executor{app: app}
  return e, nil
}

func (e *Executor) Exec() {

}
