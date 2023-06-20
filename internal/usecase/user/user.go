package user

import (
	"context"
	"t-mail/internal/entity"
	"t-mail/internal/usecase"
)

type userUseCase struct {
	creator     Creator
	authService AuthService
	finder      Finder
}

func CreateUserUseCase(creator Creator, authService AuthService, finder Finder) usecase.UserUseCase {
	return &userUseCase{creator, authService, finder}
}

func (uc *userUseCase) Register(ctx context.Context, name, email, password string) error {
	userEntity := &entity.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	if err := uc.creator.CreateUser(ctx, userEntity); err != nil {
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
	return uc.finder.Find(ctx, id)
}
