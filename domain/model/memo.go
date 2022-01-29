package model

import (
	"errors"
	"unicode/utf8"
)

type Memo struct {
	ID      int
	Content Content
}

// 新しいMemoモデルを作成します
func NewMemo(c string) (*Memo, error) {
	content, err := NewContent(c)
	if err != nil {
		return nil, err
	}
	return &Memo{ID: 0, Content: content}, nil
}

// メモの内容を編集します
func (m *Memo) SetContent(c string) error {
	replaced, err := NewContent(c)
	if err != nil {
		return err
	}
	m.Content = replaced
	return nil
}

type Content string

func NewContent(c string) (Content, error) {
	if c == "" {
		return "", errors.New("content cannot be blank")
	}
	if utf8.RuneCountInString(c) >= 500 {
		return "", errors.New("content must be less than 500 characters")
	}
	return Content(c), nil
}
