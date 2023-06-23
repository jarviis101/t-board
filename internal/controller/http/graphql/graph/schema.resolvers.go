package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"fmt"
	"t-board/internal/controller/http/graphql/directives"
	"t-board/internal/controller/http/graphql/graph/model"
)

// CreateBoard is the resolver for the createBoard field.
func (r *mutationResolver) CreateBoard(ctx context.Context, input model.CreateBoard) (*model.Board, error) {
	currentUserId := ctx.Value(directives.AuthKey(directives.Key)).(string)
	currentUser, err := r.userUseCase.Get(ctx, currentUserId)
	if err != nil {
		return nil, err
	}

	board, err := r.boardUseCase.Create(ctx, input.Title, input.Description, currentUser.ID, string(input.Type))
	if err != nil {
		return nil, err
	}

	if err = r.userUseCase.AddBoard(ctx, currentUser, board); err != nil {
		return nil, err
	}

	m := r.boardTransformer.TransformToModel(board)

	return m, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	currentUserId := ctx.Value(directives.AuthKey(directives.Key)).(string)
	currentUser, err := r.userUseCase.Get(ctx, currentUserId)
	if err != nil {
		return nil, err
	}

	m := r.userTransformer.TransformToModel(currentUser)

	return m, nil
}

// GetBoards is the resolver for the getBoards field.
func (r *queryResolver) GetBoards(ctx context.Context) ([]*model.Board, error) {
	currentUserId := ctx.Value(directives.AuthKey(directives.Key)).(string)
	currentUser, err := r.userUseCase.Get(ctx, currentUserId)
	if err != nil {
		return nil, err
	}

	var boards []*model.Board
	boardsEntity, err := r.boardUseCase.GetByUser(ctx, currentUser.ID)
	if err != nil {
		return nil, err
	}

	for _, board := range boardsEntity {
		m := r.boardTransformer.TransformToModel(board)

		boards = append(boards, m)
	}

	return boards, nil
}

// GetNotesByBoard is the resolver for the getNotesByBoard field.
func (r *queryResolver) GetNotesByBoard(ctx context.Context, board string) ([]*model.Note, error) {
	panic(fmt.Errorf("not implemented: GetNotesByBoard - getNotesByBoard"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
