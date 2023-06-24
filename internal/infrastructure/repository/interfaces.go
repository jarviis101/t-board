package repository

import (
	"context"
	"t-board/internal/entity"
)

type (
	UserRepository interface {
		Store(ctx context.Context, user *entity.User) error
		AddBoard(ctx context.Context, u *entity.User, b *entity.Board) error
		DeleteBoard(ctx context.Context, board string) error
		GetByEmail(ctx context.Context, email string) (*entity.User, error)
		GetById(ctx context.Context, id string) (*entity.User, error)
		GetByIds(ctx context.Context, ids []string) ([]*entity.User, error)
	}
	BoardRepository interface {
		Store(ctx context.Context, b *entity.Board) (*entity.Board, error)
		AddUser(ctx context.Context, u *entity.User, b *entity.Board) error
		Clear(ctx context.Context, board string) error
		Delete(ctx context.Context, board string) error
		GetByUser(ctx context.Context, user string) ([]*entity.Board, error)
		GetById(ctx context.Context, id string) (*entity.Board, error)
	}
	NoteRepository interface {
	}
)
