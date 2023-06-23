package mapper

import (
	"t-board/internal/entity"
	"t-board/internal/infrastructure/repository/mongo"
)

type UserMapper interface {
	SchemaToEntity(user *mongo.User) *entity.User
}

type userMapper struct {
	BaseMapper
}

func CreateUserMapper(bm BaseMapper) UserMapper {
	return &userMapper{bm}
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
