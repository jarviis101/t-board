package transformers

import (
	"t-board/internal/controller/http/graphql/graph/model"
	"t-board/internal/entity"
)

type UserTransformer interface {
	TransformToModel(u *entity.User) *model.User
	TransformManyToModel(u []entity.User) []*model.User
}

type userTransformer struct {
}

func CreateUserTransformer() UserTransformer {
	return &userTransformer{}
}

func (t *userTransformer) TransformManyToModel(u []entity.User) []*model.User {
	var users []*model.User
	for _, user := range u {
		m := t.TransformToModel(&user)

		users = append(users, m)
	}

	return users
}

func (t *userTransformer) TransformToModel(u *entity.User) *model.User {
	return &model.User{ID: u.ID, Email: u.Email, Name: u.Name}
}
