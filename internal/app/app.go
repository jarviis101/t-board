package app

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"t-mail/internal/controller/http"
	http_validator "t-mail/internal/controller/http/validator"
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
	userFinder := user.CreateFinder(userRepository)
	userUseCase := user.CreateUserUseCase(userCreator, userAuthService, userFinder)
	httpValidator := http_validator.CreateValidator(validator.New())

	http.RunServer(userUseCase, httpValidator)

	return nil
}
