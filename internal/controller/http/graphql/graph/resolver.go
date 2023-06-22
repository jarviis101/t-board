package graph

import (
	"t-board/internal/controller/http/graphql/transformers"
	"t-board/internal/usecase"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userUseCase      usecase.UserUseCase
	boardUseCase     usecase.BoardUseCase
	userTransformer  transformers.UserTransformer
	boardTransformer transformers.BoardTransformer
}

func CreateResolver(
	u usecase.UserUseCase,
	b usecase.BoardUseCase,
	ut transformers.UserTransformer,
	bt transformers.BoardTransformer,
) *Resolver {
	return &Resolver{u, b, ut, bt}
}
