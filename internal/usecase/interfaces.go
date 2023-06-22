package usecase

import (
	"context"
	"t-board/internal/entity"
)

type (
	UserUseCase interface {
		Register(ctx context.Context, name, email, password string) error
		Login(ctx context.Context, email, password string) (string, error)
		Get(ctx context.Context, id string) (*entity.User, error)
		GetMany(ctx context.Context, ids []string) ([]entity.User, error)
	}

	BoardUseCase interface {
		Create(ctx context.Context, title, description, creator, boardType string) (*entity.Board, error)
		GetByUser(ctx context.Context, creator string) ([]entity.Board, error)
		Clear(ctx context.Context, board string) error
		Delete(ctx context.Context, board string) error
	}

	NoteUseCase interface {
		Create(ctx context.Context, board, description string) error
		GetByBoard(ctx context.Context, board string) []interface{}
		Delete(ctx context.Context, item string) error
	}
)
