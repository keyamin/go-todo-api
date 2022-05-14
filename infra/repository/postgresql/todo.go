package postgresql

import (
	"go-todo-api/domain/model"
	"go-todo-api/domain/repository"

	"github.com/jinzhu/gorm"
)

type TodoRepositoryImpl struct {
	Conn *gorm.DB
}

type (
	Todo struct {
		ID  int `gorm:"primaryKey"`
		Title string
		IsDone bool
	}
)

func NewTodoRepository(c *gorm.DB) repository.TodoRepository {
	return &TodoRepositoryImpl{Conn: c}
}

func (tr *TodoRepositoryImpl) Create(todo *model.Todo) (*model.Todo, error) {
	entity := &Todo{
		ID:  todo.ID,
		Title: string(todo.Title),
		IsDone: todo.IsDone,
	}
	if err := tr.Conn.Table("todos").Create(entity).Error; err != nil {
		return nil, err
	}

	return todo, nil
}
