package server

import (
	"github.com/go-playground/validator/v10"
	"t-board/internal/app"
	"t-board/internal/app/di"
	"t-board/internal/controller/http"
	http_validator "t-board/internal/controller/http/validator"
	"t-board/pkg"
)

type application struct {
	container    di.Container
	serverConfig pkg.Server
}

func CreateApplication(c di.Container, sc pkg.Server) app.Application {
	return &application{c, sc}
}

func (a *application) Run() error {
	boardUseCase := a.container.ProvideBoardUseCase()
	userUseCase := a.container.ProvideUserUseCase()
	httpValidator := http_validator.CreateValidator(validator.New())
	httpServer := http.CreateServer(a.serverConfig, httpValidator, userUseCase, boardUseCase)

	return httpServer.RunServer()
}
