package board

import (
	"context"
	"t-board/internal/entity"
	"t-board/internal/infrastructure/repository"
)

type Finder interface {
	Find(ctx context.Context, id string) (*entity.Board, error)
	FindByUser(ctx context.Context, id string) ([]*entity.Board, error)
	FindOneByOwner(ctx context.Context, board, user string) (*entity.Board, error)
}

type finder struct {
	repository repository.BoardRepository
}

func CreateFinder(r repository.BoardRepository) Finder {
	return &finder{r}
}

func (f *finder) Find(ctx context.Context, id string) (*entity.Board, error) {
	return f.repository.GetById(ctx, id)
}

func (f *finder) FindOneByOwner(ctx context.Context, board, user string) (*entity.Board, error) {
	return f.repository.GetOneByOwner(ctx, board, user)
}

func (f *finder) FindByUser(ctx context.Context, user string) ([]*entity.Board, error) {
	return f.repository.GetByUser(ctx, user)
}
