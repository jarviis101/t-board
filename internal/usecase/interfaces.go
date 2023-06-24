package usecase

import (
	"context"
	"t-board/internal/entity"
)

type (
	UserUseCase interface {
		Register(ctx context.Context, user *entity.User) error
		Login(ctx context.Context, email, password string) (string, error)
		AddBoard(ctx context.Context, user *entity.User, board *entity.Board) error
		DeleteBoardFromUsers(ctx context.Context, board string) error
		Get(ctx context.Context, id string) (*entity.User, error)
		GetMany(ctx context.Context, ids []string) ([]*entity.User, error)
	}

	BoardUseCase interface {
		Create(ctx context.Context, board *entity.Board) (*entity.Board, error)
		Clear(ctx context.Context, board string) error
		Delete(ctx context.Context, board string) error
		AddUser(ctx context.Context, user *entity.User, board *entity.Board) error
		Get(ctx context.Context, id string) (*entity.Board, error)
		GetOneByOwner(ctx context.Context, board, user string) (*entity.Board, error)
		GetByUser(ctx context.Context, creator string) ([]*entity.Board, error)
	}

	NoteUseCase interface {
		Create(ctx context.Context, board, description string) error
		GetByBoard(ctx context.Context, board string) []interface{}
		Delete(ctx context.Context, item string) error
	}
)
