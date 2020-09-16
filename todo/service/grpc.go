//go:generate protoc ./todo.proto --go_out=plugins=grpc:./pb
package service

import (
	"context"
	"fmt"
	"net"

	"github.com/ks6088ts/spidey/todo/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterTodoServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostTodo(ctx context.Context, r *pb.PostTodoRequest) (*pb.PostTodoResponse, error) {
	a, err := s.service.PostTodo(ctx, r.Name)
	if err != nil {
		return nil, err
	}
	return &pb.PostTodoResponse{Todo: &pb.Todo{
		Id:   a.ID,
		Name: a.Name,
	}}, nil
}

func (s *grpcServer) GetTodo(ctx context.Context, r *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	a, err := s.service.GetTodo(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetTodoResponse{
		Todo: &pb.Todo{
			Id:   a.ID,
			Name: a.Name,
		},
	}, nil
}

func (s *grpcServer) GetTodos(ctx context.Context, r *pb.GetTodosRequest) (*pb.GetTodosResponse, error) {
	res, err := s.service.GetTodos(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	todos := []*pb.Todo{}
	for _, p := range res {
		todos = append(
			todos,
			&pb.Todo{
				Id:   p.ID,
				Name: p.Name,
			},
		)
	}
	return &pb.GetTodosResponse{Todos: todos}, nil
}
