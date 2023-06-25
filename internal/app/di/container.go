package di

import (
	"go.mongodb.org/mongo-driver/mongo"
	"t-board/internal/infrastructure/repository/mongo/mapper"
	"t-board/internal/infrastructure/repository/mongo/repository"
	"t-board/internal/pkg/hasher"
	"t-board/internal/pkg/jwt"
	"t-board/internal/usecase"
	"t-board/internal/usecase/board"
	"t-board/internal/usecase/user"
	"t-board/pkg"
)

type Container interface {
	ProvideUserUseCase() usecase.UserUseCase
	ProvideBoardUseCase() usecase.BoardUseCase
}

type container struct {
	baseRepository repository.BaseRepository
	baseMapper     mapper.BaseMapper
	database       *mongo.Database
	serverConfig   pkg.Server
}

func CreateContainer(db *mongo.Database, sc pkg.Server) Container {
	br := repository.CreateBaseRepository()
	bm := mapper.CreateBaseMapper()

	return &container{br, bm, db, sc}
}

func (c *container) ProvideUserUseCase() usecase.UserUseCase {
	return c.resolveUserUseCaseDependencies(c.baseRepository, c.baseMapper)
}

func (c *container) ProvideBoardUseCase() usecase.BoardUseCase {
	return c.resolveBoardUseCaseDependencies(c.baseRepository, c.baseMapper)
}

func (c *container) resolveUserUseCaseDependencies(
	br repository.BaseRepository,
	bm mapper.BaseMapper,
) usecase.UserUseCase {
	h := hasher.CreateManager()
	jwtManager := jwt.CreateManager(c.serverConfig.Secret)

	userMapper := mapper.CreateUserMapper(bm)
	userRepository := repository.CreateUserRepository(br, c.database.Collection("users"), userMapper)
	userCreator := user.CreateCreator(userRepository, h)
	userAuthService := user.CreateAuthService(userRepository, h, jwtManager)
	userFinder := user.CreateFinder(userRepository)
	userCollector := user.CreateCollector(userRepository)

	return user.CreateUserUseCase(userCreator, userAuthService, userFinder, userCollector)
}

func (c *container) resolveBoardUseCaseDependencies(
	br repository.BaseRepository,
	bm mapper.BaseMapper,
) usecase.BoardUseCase {
	boardMapper := mapper.CreateBoardMapper(bm)
	boardRepository := repository.CreateBoardRepository(br, c.database.Collection("boards"), boardMapper)
	boardCreator := board.CreateCreator(boardRepository)
	boardFinder := board.CreateFinder(boardRepository)
	boardCollector := board.CreateCollector(boardRepository)

	return board.CreateBoardUseCase(boardCreator, boardFinder, boardCollector)
}
