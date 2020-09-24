package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ks6088ts/spidey/api/domains"
	"github.com/ks6088ts/spidey/todo/service"
)

func NewGrpcTodoRepository(url string) (TodoRepository, error) {
	client, err := service.NewClient(url)
	if err != nil {
		return nil, err
	}

	return &grpcTodoRepository{
		client: client,
	}, nil
}

var _ TodoRepository = (*grpcTodoRepository)(nil)

type grpcTodoRepository struct {
	client *service.Client
}

func (r *grpcTodoRepository) Get(ctx context.Context, id domains.TodoID) (*domains.Todo, error) {
	fmt.Println("not implemented")
	return nil, ErrBadRequest
}

func (r *grpcTodoRepository) GetAll(ctx context.Context) ([]*domains.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	todoList, err := r.client.GetTodos(ctx, 0, 100)
	if err != nil {
		return nil, err
	}

	var todos []*domains.Todo
	for _, t := range todoList {
		todo := &domains.Todo{
			ID:        domains.TodoID(t.ID),
			Name:      t.Name,
			CreatedAt: time.Now(),
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *grpcTodoRepository) Create(ctx context.Context, todo *domains.Todo) (*domains.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	a, err := r.client.PostTodo(ctx, todo.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &domains.Todo{
		ID:        domains.TodoID(a.ID),
		Name:      a.Name,
		CreatedAt: time.Now(),
	}, nil
}

func (repo *grpcTodoRepository) Update(ctx context.Context, todo *domains.Todo) (*domains.Todo, error) {
	fmt.Println("not implemented")
	return nil, ErrBadRequest
}
