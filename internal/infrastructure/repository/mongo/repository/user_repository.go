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

type userRepository struct {
	collection *mongo.Collection
	mapper     mapper.UserMapper
}

func CreateUserRepository(c *mongo.Collection, m mapper.UserMapper) repository.UserRepository {
	return &userRepository{c, m}
}

func (r *userRepository) Store(ctx context.Context, u *entity.User) error {
	user := &schema.User{
		ID:        primitive.NewObjectID(),
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := r.collection.InsertOne(ctx, user); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user *schema.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return r.mapper.SchemaToEntity(user), nil
}

func (r *userRepository) GetById(ctx context.Context, id string) (*entity.User, error) {
	user, err := r.getById(ctx, id)
	if err != nil {
		return nil, err
	}

	return r.mapper.SchemaToEntity(user), nil
}

func (r *userRepository) GetByIds(ctx context.Context, ids []string) ([]*entity.User, error) {
	var users []*schema.User
	var usersEntity []*entity.User
	objectIds := r.fromStringToObjectId(ids)
	cur, err := r.collection.Find(ctx, bson.M{"_id": bson.M{
		"$in": objectIds,
	}})
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var user *schema.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	for _, u := range users {
		user := r.mapper.SchemaToEntity(u)
		usersEntity = append(usersEntity, user)
	}

	return usersEntity, nil
}

func (r *userRepository) getById(ctx context.Context, id string) (*schema.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user *schema.User
	err = r.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) fromStringToObjectId(ids []string) []primitive.ObjectID {
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
