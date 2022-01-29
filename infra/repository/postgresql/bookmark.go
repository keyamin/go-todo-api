package postgresql

import (
	"bookshelf/domain/model"
	"bookshelf/domain/repository"

	"github.com/jinzhu/gorm"
)

type BookmarkRepositoryImpl struct {
	Conn *gorm.DB
}

type (
	Bookmark struct {
		ID  int `gorm:"primaryKey"`
		Url string
	}
)

func NewBookmarkRepository(c *gorm.DB) repository.BookmarkRepository {
	return &BookmarkRepositoryImpl{Conn: c}
}

func (br *BookmarkRepositoryImpl) Create(bookmark *model.Bookmark) (*model.Bookmark, error) {
	entity := &Bookmark{
		ID:  bookmark.ID,
		Url: string(bookmark.Url),
	}
	if err := br.Conn.Table("bookmarks").Create(entity).Error; err != nil {
		return nil, err
	}

	return bookmark, nil
}
