package user

import (
	"context"
	"t-board/internal/entity"
	"t-board/internal/infrastructure/repository"
)

type Collector interface {
	AddBoardToUser(ctx context.Context, u *entity.User, b *entity.Board) error
	DeleteBoardFromUsers(ctx context.Context, board string) error
}

type collector struct {
	repository repository.UserRepository
}

func CreateCollector(r repository.UserRepository) Collector {
	return &collector{r}
}

func (c *collector) AddBoardToUser(ctx context.Context, u *entity.User, b *entity.Board) error {
	return c.repository.AddBoard(ctx, u, b)
}

func (c *collector) DeleteBoardFromUsers(ctx context.Context, board string) error {
	return c.repository.DeleteBoard(ctx, board)
}
