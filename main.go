package main

import (
	"bookshelf/config"
	"bookshelf/infra/repository/postgresql"
	"bookshelf/interface/controller"
	"bookshelf/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := postgresql.NewBookmarkRepository(config.NewDB())
	usecase := usecase.NewBookmarkUsecase(repository)
	controller := controller.NewBookmarkController(usecase)

	r := gin.Default()
	r.POST("/bookmark", controller.Post)

	r.Run(":8080")
}
