package board

import (
	"context"
	"t-board/internal/entity"
	"t-board/internal/infrastructure/repository"
)

type Finder interface {
	FindByUser(ctx context.Context, id string) ([]*entity.Board, error)
}

type finder struct {
	repository repository.BoardRepository
}

func CreateFinder(r repository.BoardRepository) Finder {
	return &finder{r}
}

func (f finder) FindByUser(ctx context.Context, user string) ([]*entity.Board, error) {
	return f.repository.GetByUser(ctx, user)
}
