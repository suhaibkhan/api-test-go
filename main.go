package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suhaibkhan/apitestgo/internal/handler"
	"github.com/suhaibkhan/apitestgo/internal/repository"
)

func main() {

	db := repository.Init()
	todoRepo := repository.NewTodoRepository(db)

	router := gin.Default()
	apiGroup := router.Group("/api")

	handler.RegisterTodoRoutes(apiGroup, todoRepo)

	router.Run(":8080")
}
