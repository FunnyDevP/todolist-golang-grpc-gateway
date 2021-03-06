package handler

import (
	"context"
	"fmt"
	"github.com/FunnyDevP/todolist-golang-grpc-gateway/api/proto"
	"github.com/FunnyDevP/todolist-golang-grpc-gateway/internal/repository"
	"github.com/golang/protobuf/ptypes/empty"
)

type todolistHandler struct {
	todoRepo repository.TodoListRepository
}

func NewTodolistHandler(todoRepo repository.TodoListRepository) *todolistHandler {
	return &todolistHandler{todoRepo: todoRepo}
}

func (h todolistHandler) Create(ctx context.Context, req *proto.CreateTaskRequest) (*proto.CreateTaskResponse, error) {
	return nil, nil
}

func (h todolistHandler) Gets(ctx context.Context, req *empty.Empty) (*proto.GetsResponse, error) {
	fmt.Println("call gets ")
	return nil, nil
}
