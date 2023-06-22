package user

import (
	"context"
	"t-board/internal/entity"
	"t-board/internal/usecase"
)

type userUseCase struct {
	creatorService Creator
	authService    AuthService
	finderService  Finder
}

func CreateUserUseCase(c Creator, a AuthService, f Finder) usecase.UserUseCase {
	return &userUseCase{c, a, f}
}

func (uc *userUseCase) Register(ctx context.Context, name, email, password string) error {
	userEntity := &entity.User{Name: name, Email: email, Password: password}
	if err := uc.creatorService.CreateUser(ctx, userEntity); err != nil {
		return err
	}

	return nil
}

func (uc *userUseCase) Login(ctx context.Context, email, password string) (string, error) {
	token, err := uc.authService.Authenticate(ctx, email, password)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *userUseCase) Get(ctx context.Context, id string) (*entity.User, error) {
	return uc.finderService.Find(ctx, id)
}

func (uc *userUseCase) GetMany(ctx context.Context, ids []string) ([]*entity.User, error) {
	return uc.finderService.FindMany(ctx, ids)
}
