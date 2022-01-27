package controller

import (
	"bookshelf/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookmarkController interface {
	Post(c *gin.Context)
}

type BookmarkControllerImpl struct {
	bookmarkUsecase usecase.BookmarkUsecase
}

func NewBookmarkController(bookmarkUsecase usecase.BookmarkUsecase) BookmarkController {
	return &BookmarkControllerImpl{bookmarkUsecase}
}

type requestBookmark struct {
	Url string `json:"url" binding:"required"`
}
type responseBookmark struct {
	ID  int    `json:"id"`
	Url string `json:"url"`
}

func (b *BookmarkControllerImpl) Post(c *gin.Context) {
	var req requestBookmark
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	createdModel, usecaseErr := b.bookmarkUsecase.Create(req.Url)
	if usecaseErr != nil {
		panic(usecaseErr)
	}
	created := &responseBookmark{
		ID:  createdModel.ID,
		Url: string(createdModel.Url),
	}

	c.JSON(http.StatusCreated, created)
}
