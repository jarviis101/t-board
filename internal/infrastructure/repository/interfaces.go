package repository

import (
	"context"
	"t-board/internal/entity"
)

type (
	UserRepository interface {
		Store(ctx context.Context, user *entity.User) error
		GetByEmail(ctx context.Context, email string) (*entity.User, error)
		GetById(ctx context.Context, id string) (*entity.User, error)
	}
)
