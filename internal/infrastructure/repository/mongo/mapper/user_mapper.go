package mapper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"t-board/internal/entity"
	"t-board/internal/infrastructure/repository/mongo"
)

type UserMapper interface {
	SchemaToEntity(user *mongo.User) *entity.User
}

type userMapper struct {
}

func CreateUserMapper() UserMapper {
	return &userMapper{}
}

func (u *userMapper) SchemaToEntity(user *mongo.User) *entity.User {
	return &entity.User{
		ID:        user.ID.Hex(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Password:  user.Password,
		Boards:    u.fromObjectIdToString(user.Boards),
	}
}

func (u *userMapper) fromObjectIdToString(objectIds []primitive.ObjectID) []string {
	var ids []string

	for _, objectId := range objectIds {
		ids = append(ids, objectId.Hex())
	}

	return ids
}
