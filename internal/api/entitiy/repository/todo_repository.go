package repository

import (
	"fmt"
	"gRPC-Todo/internal/api/entitiy/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITodoRepository interface {
	CreateTodo(todo *model.Todo) error
	GetTodo(todo *model.Todo, userId uint, todoId uint) error
	GetAllTodos(todo *[]model.Todo, userId uint) error
	UpdateTodo(todo *model.Todo, userId uint, todoId uint) error
	DeleteTodoById(userId uint, todoId uint) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) ITodoRepository {
	return &todoRepository{db}
}

func (tr *todoRepository) CreateTodo(todo *model.Todo) error {
	if err := tr.db.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) GetTodo(todo *model.Todo, userId uint, todoId uint) error {
	if err := tr.db.Where("user_id = ? AND id = ?", userId, todoId).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) GetAllTodos(todo *[]model.Todo, userId uint) error {
	if err := tr.db.Where("user_id = ?", userId).Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) UpdateTodo(todo *model.Todo, userId uint, todoId uint) error {
	result := tr.db.Model(todo).Clauses(clause.Returning{}).Where("id=? AND user_id = ?", todoId, userId).Update("title", todo.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("task not found")
	}
	return nil
}

func (tr *todoRepository) DeleteTodoById(userId uint, todoId uint) error {
	result := tr.db.Where("id=? AND user_id = ?", todoId, userId).Delete(&model.Todo{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("task not found")
	}
	return nil
}
