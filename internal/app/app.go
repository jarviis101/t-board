package app

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"t-board/internal/controller/http"
	http_validator "t-board/internal/controller/http/validator"
	"t-board/internal/infrastructure/repository/mongo/mapper"
	"t-board/internal/infrastructure/repository/mongo/repository"
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

	baseRepository := repository.CreateBaseRepository()
	baseMapper := mapper.CreateBaseMapper()
	boardUseCase := a.resolveBoardUseCase(baseRepository, baseMapper)
	userUseCase := a.resolveUserUseCase(baseRepository, baseMapper)

	http.RunServer(httpValidator, a.serverConfig, userUseCase, boardUseCase)

	return nil
}

func (a *application) resolveUserUseCase(br repository.BaseRepository, bm mapper.BaseMapper) usecase.UserUseCase {
	h := hasher.CreateManager()
	jwtManager := jwt.CreateManager(a.serverConfig.Secret)

	userMapper := mapper.CreateUserMapper(bm)
	userRepository := repository.CreateUserRepository(br, a.database.Collection("users"), userMapper)
	userCreator := user.CreateCreator(userRepository, h)
	userAuthService := user.CreateAuthService(userRepository, h, jwtManager)
	userFinder := user.CreateFinder(userRepository)
	userCollector := user.CreateCollector(userRepository)

	return user.CreateUserUseCase(userCreator, userAuthService, userFinder, userCollector)
}

func (a *application) resolveBoardUseCase(br repository.BaseRepository, bm mapper.BaseMapper) usecase.BoardUseCase {
	boardMapper := mapper.CreateBoardMapper(bm)
	boardRepository := repository.CreateBoardRepository(br, a.database.Collection("boards"), boardMapper)
	boardCreator := board.CreateCreator(boardRepository)
	boardFinder := board.CreateFinder(boardRepository)

	return board.CreateBoardUseCase(boardCreator, boardFinder)
}
