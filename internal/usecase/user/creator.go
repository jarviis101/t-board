package user

import (
	"context"
	"t-board/internal/entity"
	"t-board/internal/infrastructure/repository"
	"t-board/internal/pkg/hasher"
)

type Creator interface {
	CreateUser(ctx context.Context, user *entity.User) error
}

type creator struct {
	repository repository.UserRepository
	hasher     hasher.Manager
}

func CreateCreator(r repository.UserRepository, h hasher.Manager) Creator {
	return &creator{r, h}
}

func (c *creator) CreateUser(ctx context.Context, user *entity.User) error {
	user.Password = c.hasher.HashPassword(user.Password)
	if err := c.repository.Store(ctx, user); err != nil {
		return err
	}

	return nil
}
