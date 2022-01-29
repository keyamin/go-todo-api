package model

import (
	"errors"
)

type Bookmark struct {
	ID  int
	Url Url
}

// 新しいBookmarkモデルを生成します
func NewBookmark(url string) (*Bookmark, error) {
	urlObj, err := NewUrl(url)
	if err != nil {
		return nil, err
	}
	return &Bookmark{ID: 0, Url: urlObj}, nil
}

// Urlを編集します
func (b *Bookmark) SetUrl(url string) error {
	replaced, err := NewUrl(url)
	if err != nil {
		return err
	}
	b.Url = replaced
	return nil
}

type Url string

func NewUrl(url string) (Url, error) {
	if url == "" {
		return "", errors.New("URL cannot be blank")
	}

	return Url(url), nil
}
