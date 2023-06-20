package repository

import (
	"context"
	"t-mail/internal/entity"
)

type (
	UserRepository interface {
		Store(ctx context.Context, user *entity.User) error
	}
)
