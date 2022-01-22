package repository

import (
	"bookshelf/domain/model"
)

type BookmarkRepository interface {
	Create(*model.Bookmark) (*model.Bookmark, error)
}
