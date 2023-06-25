package board

import (
	"context"
	"t-board/internal/entity"
	"t-board/internal/usecase"
)

type boardUseCase struct {
	creatorService   Creator
	finderService    Finder
	collectorService Collector
}

func CreateBoardUseCase(c Creator, f Finder, cs Collector) usecase.BoardUseCase {
	return &boardUseCase{c, f, cs}
}

func (bc *boardUseCase) Create(ctx context.Context, board *entity.Board) (*entity.Board, error) {
	return bc.creatorService.CreateBoard(ctx, board)
}

func (bc *boardUseCase) Clear(ctx context.Context, board string) error {
	return nil
}

func (bc *boardUseCase) Delete(ctx context.Context, board string) error {
	return bc.collectorService.DeleteBoard(ctx, board)
}

func (bc *boardUseCase) AddUser(ctx context.Context, user *entity.User, board *entity.Board) error {
	return bc.collectorService.AddUserToBoard(ctx, user, board)
}

func (bc *boardUseCase) Get(ctx context.Context, id string) (*entity.Board, error) {
	return bc.finderService.Find(ctx, id)
}

func (bc *boardUseCase) GetOneByOwner(ctx context.Context, board, user string) (*entity.Board, error) {
	return bc.finderService.FindOneByOwner(ctx, board, user)
}

func (bc *boardUseCase) GetByUser(ctx context.Context, creator string) ([]*entity.Board, error) {
	return bc.finderService.FindByUser(ctx, creator)
}
