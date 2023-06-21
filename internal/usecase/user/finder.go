package user

import (
	"context"
	"t-board/internal/entity"
	"t-board/internal/infrastructure/repository"
)

type Finder interface {
	Find(ctx context.Context, id string) (*entity.User, error)
}

type finder struct {
	repository repository.UserRepository
}

func CreateFinder(repository repository.UserRepository) Finder {
	return &finder{repository}
}

func (f *finder) Find(ctx context.Context, id string) (*entity.User, error) {
	return f.repository.GetById(ctx, id)
}
