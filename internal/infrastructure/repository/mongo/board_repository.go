package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"t-board/internal/entity"
	"t-board/internal/infrastructure/repository"
	"time"
)

type boardRepository struct {
	collection *mongo.Collection
}

func CreateBoardRepository(c *mongo.Collection) repository.BoardRepository {
	return &boardRepository{c}
}

func (r *boardRepository) Store(ctx context.Context, b *entity.Board) (*entity.Board, error) {
	board := &Board{
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

	return b, nil
}

func (r *boardRepository) GetByUser(ctx context.Context, user string) ([]entity.Board, error) {
	userObjectId, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		return nil, err
	}

	var boards []Board
	var boardsEntity []entity.Board
	cur, err := r.collection.Find(ctx, bson.M{"members": bson.M{
		"$in": []primitive.ObjectID{userObjectId},
	}})
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var board Board
		if err := cur.Decode(&board); err != nil {
			return nil, err
		}

		boards = append(boards, board)
	}

	for _, b := range boards {
		board := entity.Board{
			ID:          b.ID.Hex(),
			Title:       b.Title,
			Description: b.Description,
			Type:        entity.BoardType(b.Type),
			CreatedAt:   b.CreatedAt,
			UpdatedAt:   b.UpdatedAt,
			Members:     r.fromObjectIdToString(b.Members),
		}
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

func (r *boardRepository) fromStringToObjectId(ids []string) []primitive.ObjectID {
	var objectIds []primitive.ObjectID
	for _, objectId := range ids {
		objectId, err := primitive.ObjectIDFromHex(objectId)
		if err != nil {
			continue
		}
		objectIds = append(objectIds, objectId)
	}

	return objectIds
}

func (r *boardRepository) fromObjectIdToString(objectIds []primitive.ObjectID) []string {
	var ids []string
	for _, objectId := range objectIds {
		ids = append(ids, objectId.Hex())
	}

	return ids
}
