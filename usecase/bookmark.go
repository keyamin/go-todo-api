package usecase

import (
	"bookshelf/domain/model"
	"bookshelf/domain/repository"
)

type BookmarkUsecase interface {
	Create(url string) (*model.Bookmark, error)
}

type BookmarkUsecaseImpl struct {
	bookmarkRepo repository.BookmarkRepository
}

func NewBookmarkUsecase(bookmarkRepo repository.BookmarkRepository) *BookmarkUsecaseImpl {
	return &BookmarkUsecaseImpl{bookmarkRepo}
}

func (b *BookmarkUsecaseImpl) Create(url string) (*model.Bookmark, error) {
	bookmark, err := model.NewBookmark(url)
	if err != nil {
		return nil, err
	}

	result, err := b.bookmarkRepo.Create(bookmark)
	if err != nil {
		return nil, err
	}

	return result, nil
}
