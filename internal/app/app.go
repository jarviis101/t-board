package app

import (
	"go.mongodb.org/mongo-driver/mongo"
	"t-mail/internal/controller/http"
	repo "t-mail/internal/infrastructure/repository/mongo"
	"t-mail/internal/pkg/hasher"
	"t-mail/internal/usecase/user"
)

type Application interface {
	Run() error
}

type application struct {
	database *mongo.Database
}

func CreateApplication(database *mongo.Database) Application {
	return &application{database}
}

func (a *application) Run() error {
	m := hasher.CreateManager()
	userRepository := repo.CreateUserRepository(a.database.Collection("users"))
	userCreator := user.CreateCreator(userRepository, m)
	userUseCase := user.CreateUserUseCase(userCreator)
	http.RunServer(userUseCase)

	return nil
}
