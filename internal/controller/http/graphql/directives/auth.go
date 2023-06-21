package directives

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"t-mail/internal/controller/http/middleware"
)

type AuthKey string

const Key = "auth"

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	claims := middleware.GetClaims()
	if claims == nil {
		err := &gqlerror.Error{Message: "Access Denied"}
		return nil, err
	}

	c := context.WithValue(ctx, AuthKey(Key), claims.UserId)
	return next(c)
}
