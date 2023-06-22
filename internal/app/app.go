package app

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"t-board/internal/controller/http"
	http_validator "t-board/internal/controller/http/validator"
	repo "t-board/internal/infrastructure/repository/mongo"
	"t-board/internal/pkg/hasher"
	"t-board/internal/pkg/jwt"
	"t-board/internal/usecase"
	"t-board/internal/usecase/board"
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

func CreateApplication(d *mongo.Database, sc pkg.Server) Application {
	return &application{d, sc}
}

func (a *application) Run() error {
	httpValidator := http_validator.CreateValidator(validator.New())
	boardUseCase := a.resolveBoardUseCase()
	userUseCase := a.resolveUserUseCase()

	http.RunServer(httpValidator, a.serverConfig, userUseCase, boardUseCase)

	return nil
}

func (a *application) resolveUserUseCase() usecase.UserUseCase {
	h := hasher.CreateManager()
	jwtManager := jwt.CreateManager(a.serverConfig.Secret)

	userRepository := repo.CreateUserRepository(a.database.Collection("users"))
	userCreator := user.CreateCreator(userRepository, h)
	userAuthService := user.CreateAuthService(userRepository, h, jwtManager)
	userFinder := user.CreateFinder(userRepository)

	return user.CreateUserUseCase(userCreator, userAuthService, userFinder)
}

func (a *application) resolveBoardUseCase() usecase.BoardUseCase {
	boardRepository := repo.CreateBoardRepository(a.database.Collection("boards"))
	boardCreator := board.CreateCreator(boardRepository)
	boardFinder := board.CreateFinder(boardRepository)

	return board.CreateBoardUseCase(boardCreator, boardFinder)
}
