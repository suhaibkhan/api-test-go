package repository

import (
	"gorm.io/gorm"

	"github.com/suhaibkhan/apitestgo/internal/domain"
)

// https://medium.com/@wahyubagus1910/build-scalable-restful-api-with-golang-gin-gonic-framework-43793c730d10

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (repo *TodoRepository) SaveTodo(todo *domain.Todo) error {
	err := repo.DB.Save(todo).Error
	return err
}

func (repo *TodoRepository) DeleteTodo(id uint64) error {
	todo := domain.Todo{Id: id}
	err := repo.DB.Delete(&todo).Error
	return err
}

func (repo *TodoRepository) FetchTodo(id uint64) (domain.Todo, error) {
	todo := domain.Todo{Id: id}
	if err := repo.DB.First(&todo).Error; err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

func (repo *TodoRepository) FetchAllTodos() ([]domain.Todo, error) {
	var todos []domain.Todo
	if err := repo.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
