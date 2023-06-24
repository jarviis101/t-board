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
	if err := c.repository.AddUser(ctx, u, b); err != nil {
		return err
	}

	return nil
}

func (c *collector) DeleteBoard(ctx context.Context, board string) error {
	if err := c.repository.Delete(ctx, board); err != nil {
		return err
	}

	return nil
}
