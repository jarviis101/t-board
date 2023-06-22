package board

import (
	"context"
	"t-board/internal/entity"
	"t-board/internal/infrastructure/repository"
)

type Creator interface {
	CreateBoard(ctx context.Context, board *entity.Board) (*entity.Board, error)
}

type creator struct {
	repository repository.BoardRepository
}

func CreateCreator(r repository.BoardRepository) Creator {
	return &creator{r}
}

func (c *creator) CreateBoard(ctx context.Context, board *entity.Board) (*entity.Board, error) {
	b, err := c.repository.Store(ctx, board)
	if err != nil {
		return nil, err
	}

	return b, nil
}
