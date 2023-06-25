package user

import (
	"context"
	"t-board/internal/entity"
	"t-board/internal/usecase"
)

type userUseCase struct {
	creatorService   Creator
	authService      AuthService
	finderService    Finder
	collectorService Collector
}

func CreateUserUseCase(c Creator, a AuthService, f Finder, cs Collector) usecase.UserUseCase {
	return &userUseCase{c, a, f, cs}
}

func (uc *userUseCase) Register(ctx context.Context, user *entity.User) error {
	return uc.creatorService.CreateUser(ctx, user)
}

func (uc *userUseCase) Login(ctx context.Context, email, password string) (string, error) {
	return uc.authService.Authenticate(ctx, email, password)
}

func (uc *userUseCase) AddBoard(ctx context.Context, user *entity.User, board *entity.Board) error {
	return uc.collectorService.AddBoardToUser(ctx, user, board)
}

func (uc *userUseCase) DeleteBoardFromUsers(ctx context.Context, board string) error {
	return uc.collectorService.DeleteBoardFromUsers(ctx, board)
}

func (uc *userUseCase) Get(ctx context.Context, id string) (*entity.User, error) {
	return uc.finderService.Find(ctx, id)
}

func (uc *userUseCase) GetMany(ctx context.Context, ids []string) ([]*entity.User, error) {
	return uc.finderService.FindMany(ctx, ids)
}
