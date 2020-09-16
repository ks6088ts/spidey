package service

import (
	"context"

	"github.com/ks6088ts/spidey/todo/model"
	"github.com/ks6088ts/spidey/todo/repository"
	"github.com/segmentio/ksuid"
)

type Service interface {
	PostTodo(ctx context.Context, name string) (*model.Todo, error)
	GetTodo(ctx context.Context, id string) (*model.Todo, error)
	GetTodos(ctx context.Context, skip uint64, take uint64) ([]model.Todo, error)
}

type todoService struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &todoService{r}
}

func (s *todoService) PostTodo(ctx context.Context, name string) (*model.Todo, error) {
	a := &model.Todo{
		Name: name,
		ID:   ksuid.New().String(),
	}
	if err := s.repository.PutTodo(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *todoService) GetTodo(ctx context.Context, id string) (*model.Todo, error) {
	return s.repository.GetTodoByID(ctx, id)
}

func (s *todoService) GetTodos(ctx context.Context, skip uint64, take uint64) ([]model.Todo, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	return s.repository.ListTodos(ctx, skip, take)
}
