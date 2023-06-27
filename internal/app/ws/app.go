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
	boardUseCase := a.container.ProvideBoardUseCase()
	userUseCase := a.container.ProvideUserUseCase()
	wsServer := ws.CreateWSServer(userUseCase, boardUseCase)

	return wsServer.RunServer()
}
