package main

import (
	"bookshelf/config"
	"bookshelf/infrastructure"
	"bookshelf/interface/controller"
	"bookshelf/usecase"
	"net/http"
)

func main() {
	repository := infrastructure.NewBookmarkRepository(config.NewDB())
	usecase := usecase.NewBookmarkUsecase(repository)
	controller := controller.NewBookmarkController(usecase)

	http.HandleFunc("/bookmark", controller.Post())
	http.ListenAndServe(":8080", nil)
}
