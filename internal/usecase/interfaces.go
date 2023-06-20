package usecase

import "context"

type (
	UserUseCase interface {
		Register(ctx context.Context, name, email, password string) error
		Login(ctx context.Context, email, password string) (string, error)
	}
)
