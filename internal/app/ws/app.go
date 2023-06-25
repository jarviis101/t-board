package ws

import (
	"t-board/internal/app"
	"t-board/internal/app/di"
	"t-board/internal/controller/ws"
)

type application struct {
	container di.Container
}

func CreateApplication(c di.Container) app.Application {
	return &application{c}
}

func (a *application) Run() error {
	wsServer := ws.CreateWSServer()

	return wsServer.RunServer()
}
