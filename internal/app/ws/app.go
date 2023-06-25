package ws

import (
	"t-board/internal/app"
	"t-board/internal/controller/ws"
)

type application struct {
	//container di.Container
}

func CreateApplication() app.Application {
	return &application{}
}

func (a *application) Run() error {
	wsServer := ws.CreateWSServer()

	return wsServer.RunServer()
}
