package mapper

import (
	"t-board/internal/entity"
	"t-board/internal/infrastructure/repository/mongo"
)

type BoardMapper interface {
	SchemaToEntity(board *mongo.Board) *entity.Board
}

type boardMapper struct {
	BaseMapper
}

func CreateBoardMapper(bm BaseMapper) BoardMapper {
	return &boardMapper{bm}
}

func (b *boardMapper) SchemaToEntity(board *mongo.Board) *entity.Board {
	return &entity.Board{
		ID:          board.ID.Hex(),
		Title:       board.Title,
		Description: board.Description,
		Type:        entity.BoardType(board.Type),
		CreatedAt:   board.CreatedAt,
		UpdatedAt:   board.UpdatedAt,
		Members:     b.fromObjectIdToString(board.Members),
	}
}
