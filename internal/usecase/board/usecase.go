package board

import (
	"context"
	"t-board/internal/entity"
	"t-board/internal/usecase"
)

type boardUseCase struct {
	creatorService Creator
	finderService  Finder
}

func CreateBoardUseCase(c Creator, f Finder) usecase.BoardUseCase {
	return &boardUseCase{c, f}
}

func (bc *boardUseCase) Create(
	ctx context.Context,
	title, description, creator, boardType string,
) (*entity.Board, error) {
	board := &entity.Board{
		Title:       title,
		Description: description,
		Members:     []string{creator},
		Type:        entity.BoardType(boardType),
	}

	return bc.creatorService.CreateBoard(ctx, board)
}

func (bc *boardUseCase) Clear(ctx context.Context, board string) error {
	return nil
}

func (bc *boardUseCase) Delete(ctx context.Context, board string) error {
	return nil
}

func (bc *boardUseCase) Get(ctx context.Context, id string) (*entity.Board, error) {
	return bc.finderService.Find(ctx, id)
}

func (bc *boardUseCase) GetByUser(ctx context.Context, creator string) ([]*entity.Board, error) {
	return bc.finderService.FindByUser(ctx, creator)
}
