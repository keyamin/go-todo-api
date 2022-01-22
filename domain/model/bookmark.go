package model

import "errors"

type Bookmark struct {
	ID  int
	Url Url
}

func NewBookmark(url string) (*Bookmark, error) {
	urlObj, err := NewUrl(url)
	if err != nil {
		return nil, err
	}

	bookmark := &Bookmark{
		ID:  0,
		Url: *urlObj,
	}

	return bookmark, nil
}

func (b *Bookmark) SetTitle(url string) error {
	recreated, err := b.Url.SetUrl(url)
	if err != nil {
		return err
	}

	b.Url = *recreated

	return nil
}

// value object
type Url struct {
	url string
}

func NewUrl(url string) (*Url, error) {
	if url == "" {
		return nil, errors.New("URL cannot be blank")
	}

	return &Url{url}, nil
}

func (u *Url) GetUrl() string {
	return u.url
}

func (u *Url) SetUrl(url string) (*Url, error) {
	recreated, err := NewUrl(url)
	if err != nil {
		return nil, err
	}

	return recreated, nil
}
