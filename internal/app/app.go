package app

import (
	"go.mongodb.org/mongo-driver/mongo"
	"t-mail/internal/controller/http"
	repo "t-mail/internal/infrastructure/repository/mongo"
	"t-mail/internal/pkg/hasher"
	"t-mail/internal/pkg/jwt"
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
	h := hasher.CreateManager()
	jwtManager := jwt.CreateManager()
	userRepository := repo.CreateUserRepository(a.database.Collection("users"))
	userCreator := user.CreateCreator(userRepository, h)
	userAuthService := user.CreateAuthService(userRepository, h, jwtManager)
	userUseCase := user.CreateUserUseCase(userCreator, userAuthService)
	http.RunServer(userUseCase)

	return nil
}
