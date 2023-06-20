package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user *User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:       user.ID.Hex(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
