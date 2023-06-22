package repository

import (
	"context"
	"t-board/internal/entity"
)

type (
	UserRepository interface {
		Store(ctx context.Context, user *entity.User) error
		GetByEmail(ctx context.Context, email string) (*entity.User, error)
		GetById(ctx context.Context, id string) (*entity.User, error)
		GetByIds(ctx context.Context, ids []string) ([]*entity.User, error)
	}
	BoardRepository interface {
		Store(ctx context.Context, b *entity.Board) (*entity.Board, error)
		GetByUser(ctx context.Context, user string) ([]*entity.Board, error)
		Clear(ctx context.Context, board string) error
		Delete(ctx context.Context, board string)
	}
	NoteRepository interface {
	}
)
