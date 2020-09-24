package service

import (
	"context"

	"github.com/ks6088ts/spidey/todo/model"

	"github.com/ks6088ts/spidey/todo/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.TodoServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewTodoServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) PostTodo(ctx context.Context, name string) (*model.Todo, error) {
	r, err := c.service.PostTodo(
		ctx,
		&pb.PostTodoRequest{Name: name},
	)
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:   r.Todo.Id,
		Name: r.Todo.Name,
	}, nil
}

func (c *Client) GetTodo(ctx context.Context, id string) (*model.Todo, error) {
	r, err := c.service.GetTodo(
		ctx,
		&pb.GetTodoRequest{Id: id},
	)
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:   r.Todo.Id,
		Name: r.Todo.Name,
	}, nil
}

func (c *Client) GetTodos(ctx context.Context, skip uint64, take uint64) ([]model.Todo, error) {
	r, err := c.service.GetTodos(
		ctx,
		&pb.GetTodosRequest{
			Skip: skip,
			Take: take,
		},
	)
	if err != nil {
		return nil, err
	}
	todos := []model.Todo{}
	for _, a := range r.Todos {
		todos = append(todos, model.Todo{
			ID:   a.Id,
			Name: a.Name,
		})
	}
	return todos, nil
}
