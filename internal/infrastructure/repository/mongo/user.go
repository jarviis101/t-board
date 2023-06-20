package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"t-mail/internal/entity"
	"t-mail/internal/infrastructure/repository"
	"time"
)

type userRepository struct {
	collection *mongo.Collection
}

func CreateUserRepository(collection *mongo.Collection) repository.UserRepository {
	return &userRepository{collection: collection}
}

func (r *userRepository) Store(ctx context.Context, u *entity.User) error {
	user := &User{
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
