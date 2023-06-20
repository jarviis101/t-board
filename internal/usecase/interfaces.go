package usecase

import (
	"context"
	"t-mail/internal/entity"
)

type (
	UserUseCase interface {
		Register(ctx context.Context, name, email, password string) error
		Login(ctx context.Context, email, password string) (string, error)
		Get(ctx context.Context, id string) (*entity.User, error)
	}
)
