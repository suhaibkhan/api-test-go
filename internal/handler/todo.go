package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suhaibkhan/apitestgo/internal/domain"
	"github.com/suhaibkhan/apitestgo/internal/repository"
)

type TodoHandler struct {
	Repo *repository.TodoRepository
}

func RegisterTodoRoutes(apiGroup *gin.RouterGroup, todoRepo *repository.TodoRepository) {
	todoHandler := &TodoHandler{Repo: todoRepo}
	todoApiGroup := apiGroup.Group("/todos")
	{
		todoApiGroup.GET("/", todoHandler.FetchAllTodos)
		todoApiGroup.GET("/:todoId", todoHandler.FetchTodo)
		todoApiGroup.DELETE("/:todoId", todoHandler.DeleteTodo)
		todoApiGroup.POST("/", todoHandler.CreateTodo)
	}
}

func (handler *TodoHandler) CreateTodo(c *gin.Context) {
	todo := &domain.Todo{}
	if err := c.BindJSON(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo object"})
		return
	}

	if err := handler.Repo.SaveTodo(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func (handler *TodoHandler) DeleteTodo(c *gin.Context) {
	todoId, err := strconv.ParseUint(c.Param("todoId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo id"})
		return
	}

	if err := handler.Repo.DeleteTodo(todoId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.Status(http.StatusOK)
	}
}

func (handler *TodoHandler) FetchTodo(c *gin.Context) {
	todoId, err := strconv.ParseUint(c.Param("todoId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo id"})
		return
	}

	if todo, err := handler.Repo.FetchTodo(todoId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func (handler *TodoHandler) FetchAllTodos(c *gin.Context) {
	todos, err := handler.Repo.FetchAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todos)
	}
}
