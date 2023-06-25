package board

import (
	"context"
	"t-board/internal/entity"
	"t-board/internal/infrastructure/repository"
)

type Collector interface {
	AddUserToBoard(ctx context.Context, u *entity.User, b *entity.Board) error
	DeleteBoard(ctx context.Context, board string) error
}

type collector struct {
	repository repository.BoardRepository
}

func CreateCollector(r repository.BoardRepository) Collector {
	return &collector{r}
}

func (c *collector) AddUserToBoard(ctx context.Context, u *entity.User, b *entity.Board) error {
	return c.repository.AddUser(ctx, u, b)
}

func (c *collector) DeleteBoard(ctx context.Context, board string) error {
	return c.repository.Delete(ctx, board)
}
