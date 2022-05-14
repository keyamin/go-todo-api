package main

import (
	"go-todo-api/config"
	"go-todo-api/infra/repository/postgresql"
	"go-todo-api/interface/controller"
	"go-todo-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := postgresql.NewTodoRepository(config.NewDB())
	usecase := usecase.NewTodoUsecase(repository)
	controller := controller.NewTodoController(usecase)

	r := gin.Default()
	r.POST("/todos", controller.Post)

	r.Run(":8080")
}
