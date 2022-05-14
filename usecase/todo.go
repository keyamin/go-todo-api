package usecase

import (
	"go-todo-api/domain/model"
	"go-todo-api/domain/repository"
	"go-todo-api/domain/usecase"
)

type TodoUsecaseImpl struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) usecase.TodoUsecase {
	return &TodoUsecaseImpl{todoRepo: todoRepo}
}

func (t *TodoUsecaseImpl) Create(title string) (*model.Todo, error) {
	todo, err := model.NewTodo(title)
	if err != nil {
		return nil, err
	}

	result, err := t.todoRepo.Create(todo)
	if err != nil {
		return nil, err
	}

	return result, nil
}
