package mapper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"t-board/internal/entity"
	"t-board/internal/infrastructure/repository/mongo"
)

type BoardMapper interface {
	SchemaToEntity(board *mongo.Board) *entity.Board
}

type boardMapper struct {
}

func CreateBoardMapper() BoardMapper {
	return &boardMapper{}
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

func (b *boardMapper) fromObjectIdToString(objectIds []primitive.ObjectID) []string {
	var ids []string
	for _, objectId := range objectIds {
		ids = append(ids, objectId.Hex())
	}

	return ids
}
