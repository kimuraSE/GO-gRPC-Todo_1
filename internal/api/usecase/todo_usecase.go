package usecase

import (
	"gRPC-Todo/internal/api/entitiy/model"
	"gRPC-Todo/internal/api/entitiy/repository"
)

type ITodoUsecase interface {
	CreateTodo(todo model.Todo) (model.TodoResponse, error)
	GetTodo(userId uint, todoId uint) (model.TodoResponse, error)
	GetAllTodos(userId uint) ([]model.TodoResponse, error)
	UpdateTodo(todo model.Todo) (model.TodoResponse, error)
	DeleteTodoById(userId uint, todoId uint) (string, error)
}

type todoUsecase struct {
	tr repository.ITodoRepository
}

func NewTodoUsecase(tr repository.ITodoRepository) ITodoUsecase {
	return &todoUsecase{tr}
}

func (tu *todoUsecase) CreateTodo(todo model.Todo) (model.TodoResponse, error) {

	newTodo := model.Todo{
		Title:  todo.Title,
		UserID: todo.UserID,
	}

	if err := tu.tr.CreateTodo(&newTodo); err != nil {
		return model.TodoResponse{}, err
	}

	todoResponse := model.TodoResponse{
		Id:    newTodo.ID,
		Title: newTodo.Title,
	}

	return todoResponse, nil
}

func (tu *todoUsecase) GetTodo(userId uint, todoId uint) (model.TodoResponse, error) {

	newTodo := model.Todo{}

	if err := tu.tr.GetTodo(&newTodo, userId, todoId); err != nil {
		return model.TodoResponse{}, err
	}

	todoResponse := model.TodoResponse{
		Id:    newTodo.ID,
		Title: newTodo.Title,
	}

	return todoResponse, nil
}

func (tu *todoUsecase) GetAllTodos(userId uint) ([]model.TodoResponse, error) {

	newTodo := []model.Todo{}

	if err := tu.tr.GetAllTodos(&newTodo, userId); err != nil {
		return nil, err
	}

	todoResponse := []model.TodoResponse{}

	for _, todo := range newTodo {
		todoResponse = append(todoResponse, model.TodoResponse{
			Id:    todo.ID,
			Title: todo.Title,
		})
	}

	return todoResponse, nil
}

func (tu *todoUsecase) UpdateTodo(todo model.Todo) (model.TodoResponse, error) {

	newTodo := model.Todo{
		Title: todo.Title,
	}

	if err := tu.tr.UpdateTodo(&newTodo, todo.UserID, todo.ID); err != nil {
		return model.TodoResponse{}, err
	}

	todoResponse := model.TodoResponse{
		Id:    newTodo.ID,
		Title: newTodo.Title,
	}

	return todoResponse, nil
}

func (tu *todoUsecase) DeleteTodoById(userId uint, todoId uint) (string, error) {

	if err := tu.tr.DeleteTodoById(userId, todoId); err != nil {
		return "no", err
	}

	return "yes", nil
}
