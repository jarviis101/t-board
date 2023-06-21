package app

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"t-board/internal/controller/http"
	http_validator "t-board/internal/controller/http/validator"
	repo "t-board/internal/infrastructure/repository/mongo"
	"t-board/internal/pkg/hasher"
	"t-board/internal/pkg/jwt"
	"t-board/internal/usecase/user"
	"t-board/pkg"
)

type Application interface {
	Run() error
}

type application struct {
	database     *mongo.Database
	serverConfig pkg.Server
}

func CreateApplication(database *mongo.Database, serverConfig pkg.Server) Application {
	return &application{database, serverConfig}
}

func (a *application) Run() error {
	h := hasher.CreateManager()
	jwtManager := jwt.CreateManager(a.serverConfig.Secret)
	userRepository := repo.CreateUserRepository(a.database.Collection("users"))
	userCreator := user.CreateCreator(userRepository, h)
	userAuthService := user.CreateAuthService(userRepository, h, jwtManager)
	userFinder := user.CreateFinder(userRepository)
	userUseCase := user.CreateUserUseCase(userCreator, userAuthService, userFinder)
	httpValidator := http_validator.CreateValidator(validator.New())

	http.RunServer(userUseCase, httpValidator, a.serverConfig)

	return nil
}
