package model

import (
	"bookshelf/domain/value"
)

type Bookmark struct {
	ID  int
	Url value.Url
}

// 新しいBookmarkモデルを生成します
func NewBookmark(url string) (*Bookmark, error) {
	urlObj, err := value.NewUrl(url)
	if err != nil {
		return nil, err
	}
	return &Bookmark{ID: 0, Url: urlObj}, nil
}

// Urlを編集します
func (b *Bookmark) SetUrl(url string) error {
	replaced, err := value.NewUrl(url)
	if err != nil {
		return err
	}
	b.Url = replaced
	return nil
}
