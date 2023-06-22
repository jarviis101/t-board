package mapper

import (
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
	}
}
