package board

import (
	"context"
	"t-board/internal/entity"
	"t-board/internal/infrastructure/repository"
)

type Collector interface {
	AddUserToBoard(ctx context.Context, u *entity.User, b *entity.Board) error
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
