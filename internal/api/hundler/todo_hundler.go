package hundler

import (
	"context"
	"gRPC-Todo/internal/api/entitiy/model"
	"gRPC-Todo/internal/api/usecase"
	"gRPC-Todo/pkg/todo"
)

type TodoHundler interface {
	Create(ctx context.Context, in *todo.Todo) (*todo.TodoResponse, error)
	Read(ctx context.Context, in *todo.ReadRequest) (*todo.TodoResponse, error)
	GetAllTodoList(ctx context.Context, in *todo.TodoListRequest) (*todo.TodoListResponse, error)
	Update(ctx context.Context, in *todo.UpdateTodoRequest) (*todo.TodoResponse, error)
	Delete(ctx context.Context, in *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error)
}

type todoHundler struct {
	todo.UnimplementedTodoServiceServer
	tu usecase.ITodoUsecase
}

func NewTodoHundler(tu usecase.ITodoUsecase) todo.TodoServiceServer {
	return &todoHundler{tu: tu}
}

func (th *todoHundler) Create(ctx context.Context, in *todo.Todo) (*todo.TodoResponse, error) {

	newTodo := model.Todo{
		Title:  in.Title,
		UserID: uint(in.UserId),
	}

	todoResponse, err := th.tu.CreateTodo(newTodo)
	if err != nil {
		return &todo.TodoResponse{}, err
	}

	return &todo.TodoResponse{Id: uint32(todoResponse.Id), Title: todoResponse.Title}, nil
}

func (th *todoHundler) Read(ctx context.Context, in *todo.ReadRequest) (*todo.TodoResponse, error) {

	todoResponse, err := th.tu.GetTodo(uint(in.UserId), uint(in.TodoId))
	if err != nil {
		return &todo.TodoResponse{}, err
	}

	return &todo.TodoResponse{Id: uint32(todoResponse.Id), Title: todoResponse.Title}, nil
}

func (th *todoHundler) GetAllTodoList(ctx context.Context, in *todo.TodoListRequest) (*todo.TodoListResponse, error) {

	todoResponse, err := th.tu.GetAllTodos(uint(in.UserId))
	if err != nil {
		return &todo.TodoListResponse{}, err
	}

	todoResponseProto := []*todo.TodoResponse{}

	for _, v := range todoResponse {
		todoResponseProto = append(todoResponseProto, &todo.TodoResponse{Id: uint32(v.Id), Title: v.Title})
	}

	return &todo.TodoListResponse{Todos: todoResponseProto}, nil
}

func (th *todoHundler) Update(ctx context.Context, in *todo.UpdateTodoRequest) (*todo.TodoResponse, error) {

	newTodo := model.Todo{
		ID:     uint(in.TodoId),
		Title:  in.Title,
		UserID: uint(in.UserId),
	}

	todoResponse, err := th.tu.UpdateTodo(newTodo)
	if err != nil {
		return &todo.TodoResponse{}, err
	}

	return &todo.TodoResponse{Id: uint32(todoResponse.Id), Title: todoResponse.Title}, nil
}

func (th *todoHundler) Delete(ctx context.Context, in *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error) {
	
	message, err := th.tu.DeleteTodoById(uint(in.UserId), uint(in.TodoId))
	if err != nil {
		return &todo.DeleteTodoResponse{}, err
	}

	return &todo.DeleteTodoResponse{Message: message}, nil
}