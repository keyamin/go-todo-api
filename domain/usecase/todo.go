package usecase

import (
	"go-todo-api/domain/model"
)

type TodoUsecase interface {
	Create(title string) (*model.Todo, error)
}
