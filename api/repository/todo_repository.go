package repository

import (
	"context"
	"github.com/ks6088ts/spidey/api/domains"
)

type TodoRepository interface {
	Get(ctx context.Context, id domains.TodoID) (*domains.Todo, error)
	GetAll(ctx context.Context) ([]*domains.Todo, error)
	Create(ctx context.Context, todo *domains.Todo) (*domains.Todo, error)
	Update(ctx context.Context, todo *domains.Todo) (*domains.Todo, error)
}
