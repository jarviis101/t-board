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
		//AddBoard(ctx context.Context, user, board string)
	}

	BoardUseCase interface {
		Create(ctx context.Context, title, description, creator string) error
		GetByUser(ctx context.Context, creator string) []interface{}
		Clear(ctx context.Context, board string) error
		Delete(ctx context.Context, board string)
	}

	NoteUseCase interface {
		Create(ctx context.Context, board, description string) error
		GetByBoard(ctx context.Context, board string) []interface{}
		Delete(ctx context.Context, item string)
	}
)
