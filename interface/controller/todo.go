package controller

import (
	"go-todo-api/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	Post(c *gin.Context)
}

type TodoControllerImpl struct {
	TodoUsecase usecase.TodoUsecase
}

func NewTodoController(todoUsecase usecase.TodoUsecase) TodoController {
	return &TodoControllerImpl{todoUsecase}
}

type requestTodo struct {
	Title string `json:"title" binding:"required"`
}
type responseTodo struct {
	ID  int    `json:"id"`
	Title string `json:"title"`
	IsDone bool `json:"isDone"`
}

func (t *TodoControllerImpl) Post(c *gin.Context) {
	var req requestTodo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	createdModel, usecaseErr := t.TodoUsecase.Create(req.Title)
	if usecaseErr != nil {
		panic(usecaseErr)
	}
	created := &responseTodo{
		ID:  createdModel.ID,
		Title: string(createdModel.Title),
		IsDone: false,
	}

	c.JSON(http.StatusCreated, created)
}
