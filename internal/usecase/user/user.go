package user

import (
	"context"
	"t-mail/internal/entity"
	"t-mail/internal/usecase"
)

type userUseCase struct {
	creator Creator
}

func CreateUserUseCase(creator Creator) usecase.UserUseCase {
	return &userUseCase{creator}
}

func (u *userUseCase) Register(ctx context.Context, name, email, password string) error {
	userEntity := &entity.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	if err := u.creator.CreateUser(ctx, userEntity); err != nil {
		return err
	}
	return nil
}
