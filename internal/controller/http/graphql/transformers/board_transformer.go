package transformers

import (
	"t-board/internal/controller/http/graphql/graph/model"
	"t-board/internal/entity"
)

type BoardTransformer interface {
	TransformToModel(b *entity.Board) *model.Board
	TransformManyToModel(b []*entity.Board) []*model.Board
}

type boardTransformer struct {
	BaseTransformer
}

func CreateBoardTransformer(bt BaseTransformer) BoardTransformer {
	return &boardTransformer{bt}
}

func (t *boardTransformer) TransformManyToModel(b []*entity.Board) []*model.Board {
	var boards []*model.Board

	for _, board := range b {
		m := t.TransformToModel(board)

		boards = append(boards, m)
	}

	return boards
}

func (t *boardTransformer) TransformToModel(b *entity.Board) *model.Board {
	return &model.Board{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		CreatedAt:   b.CreatedAt.String(),
		UpdatedAt:   b.UpdatedAt.String(),
		Type:        model.BoardType(b.Type),
		Members:     t.modifyIds(b.Members),
	}
}
