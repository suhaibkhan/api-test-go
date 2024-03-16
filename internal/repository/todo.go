package repository

import (
	"fmt"

	"gorm.io/gorm"
)

// https://medium.com/@wahyubagus1910/build-scalable-restful-api-with-golang-gin-gonic-framework-43793c730d10

type Todo struct {
	Id   uint64 `json:"id" gorm:"primary_key"`
	Item string `json:"item"`
	Done bool   `json:"done"`
}

func (todo *Todo) ToString() string {
	return fmt.Sprintf("%+v\n", todo)
}

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (repo *TodoRepository) SaveTodo(todo *Todo) error {
	err := repo.DB.Save(todo).Error
	return err
}

func (repo *TodoRepository) DeleteTodo(id uint64) error {
	todo := Todo{Id: id}
	err := repo.DB.Delete(&todo).Error
	return err
}

func (repo *TodoRepository) FetchTodo(id uint64) (Todo, error) {
	todo := Todo{Id: id}
	if err := repo.DB.First(&todo).Error; err != nil {
		return Todo{}, err
	}
	return todo, nil
}

func (repo *TodoRepository) FetchAllTodos() ([]Todo, error) {
	var todos []Todo
	if err := repo.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
