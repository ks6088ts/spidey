package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ks6088ts/spidey/api/domains"
	"github.com/ks6088ts/spidey/api/graph/generated"
	"github.com/ks6088ts/spidey/api/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*domains.Todo, error) {
	return r.TodoRepository.Create(ctx, &domains.Todo{
		Name: input.Name,
	})
}

func (r *queryResolver) Todos(ctx context.Context) ([]*domains.Todo, error) {
	return r.TodoRepository.GetAll(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
