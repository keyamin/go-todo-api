package infrastructure

import (
	"bookshelf/domain/model"
	"bookshelf/domain/repository"

	"github.com/jinzhu/gorm"
)

type BookmarkRepositoryImpl struct {
	Conn *gorm.DB
}

func NewBookmarkRepository(c *gorm.DB) repository.BookmarkRepository {
	return &BookmarkRepositoryImpl{Conn: c}
}

func (br *BookmarkRepositoryImpl) Create(bookmark *model.Bookmark) (*model.Bookmark, error) {
	if err := br.Conn.Table("bookmarks").Create(&bookmark).Error; err != nil {
		return nil, err
	}

	return bookmark, nil
}
