package usecase

import (
	"bookshelf/domain/model"
)

type BookmarkUsecase interface {
	Create(url string) (*model.Bookmark, error)
}
