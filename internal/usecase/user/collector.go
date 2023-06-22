package user

import (
	"context"
	"t-board/internal/entity"
	"t-board/internal/infrastructure/repository"
)

type Collector interface {
	AddBoardForUser(ctx context.Context, u *entity.User, b *entity.Board) error
}

type collector struct {
	repository repository.UserRepository
}

func CreateCollector(r repository.UserRepository) Collector {
	return &collector{r}
}

func (c *collector) AddBoardForUser(ctx context.Context, u *entity.User, b *entity.Board) error {
	if err := c.repository.AddBoard(ctx, u, b); err != nil {
		return err
	}

	return nil
}
