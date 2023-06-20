package app

import (
	"context"
	"t-mail/internal/controller/http"
	"t-mail/internal/infrastructure/repository/mongo"
	"t-mail/internal/pkg/database"
	"t-mail/internal/pkg/hasher"
	"t-mail/internal/usecase/user"
)

type Application interface {
	Run() error
}

type application struct {
}

func CreateApplication() Application {
	return &application{}
}

func (a *application) Run() error {
	db := database.CreateDatabaseConnection(context.Background())
	m := hasher.CreateManager()
	userRepository := mongo.CreateUserRepository(db.Collection("users"))
	userCreator := user.CreateCreator(userRepository, m)
	userUseCase := user.CreateUserUseCase(userCreator)
	http.RunServer(userUseCase)

	return nil
}
