package repository

import (
	"go-todo-api/domain/model"
)

type TodoRepository interface {
	Create(*model.Todo) (*model.Todo, error)
}
