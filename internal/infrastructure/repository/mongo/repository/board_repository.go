package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"t-board/internal/entity"
	"t-board/internal/infrastructure/repository"
	schema "t-board/internal/infrastructure/repository/mongo"
	"t-board/internal/infrastructure/repository/mongo/mapper"
	"time"
)

type boardRepository struct {
	BaseRepository
	collection *mongo.Collection
	mapper     mapper.BoardMapper
}

func CreateBoardRepository(br BaseRepository, c *mongo.Collection, m mapper.BoardMapper) repository.BoardRepository {
	return &boardRepository{br, c, m}
}

func (r *boardRepository) Store(ctx context.Context, b *entity.Board) (*entity.Board, error) {
	board := &schema.Board{
		ID:          primitive.NewObjectID(),
		Title:       b.Title,
		Description: b.Description,
		Type:        string(b.Type),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Members:     r.fromStringToObjectId(b.Members),
	}

	if _, err := r.collection.InsertOne(ctx, board); err != nil {
		return nil, err
	}

	return r.mapper.SchemaToEntity(board), nil
}

func (r *boardRepository) GetByUser(ctx context.Context, user string) ([]*entity.Board, error) {
	userObjectId, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		return nil, err
	}

	var boards []*schema.Board
	var boardsEntity []*entity.Board
	cur, err := r.collection.Find(ctx, bson.M{"members": bson.M{
		"$in": []primitive.ObjectID{userObjectId},
	}})
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var board *schema.Board
		if err := cur.Decode(&board); err != nil {
			return nil, err
		}

		boards = append(boards, board)
	}

	for _, b := range boards {
		board := r.mapper.SchemaToEntity(b)
		boardsEntity = append(boardsEntity, board)
	}

	return boardsEntity, nil
}

func (r *boardRepository) Clear(ctx context.Context, board string) error {
	//TODO implement me
	panic("implement me")
}

func (r *boardRepository) Delete(ctx context.Context, board string) {
	//TODO implement me
	panic("implement me")
}

func (r *boardRepository) GetById(ctx context.Context, id string) (*entity.Board, error) {
	board, err := r.getById(ctx, id)
	if err != nil {
		return nil, err
	}

	return r.mapper.SchemaToEntity(board), nil
}

func (r *boardRepository) getById(ctx context.Context, id string) (*schema.Board, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var board *schema.Board
	err = r.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&board)
	if err != nil {
		return nil, err
	}

	return board, nil
}
