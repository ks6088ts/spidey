package repository

import (
	"context"

	"github.com/ks6088ts/spidey/todo/model"
)

type Repository interface {
	Close()
	PutTodo(ctx context.Context, a model.Todo) error
	GetTodoByID(ctx context.Context, id string) (*model.Todo, error)
	ListTodos(ctx context.Context, skip uint64, take uint64) ([]model.Todo, error)
}
