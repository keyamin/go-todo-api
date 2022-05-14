package model

import (
	"errors"
)

type Todo struct {
	ID  int
	Title Title
	IsDone bool
}

// 新しいBookmarkモデルを生成します
func NewTodo(title string) (*Todo, error) {
	titleObj, err := NewTitle(title)
	if err != nil {
		return nil, err
	}
	return &Todo{ID: 0, Title:titleObj,IsDone: false}, nil
}

// Urlを編集します
func (t *Todo) SetUrl(title string) error {
	replaced, err := NewTitle(title)
	if err != nil {
		return err
	}
	t.Title = replaced
	return nil
}

type Title string

func NewTitle(title string) (Title, error) {
	if title == "" {
		return "", errors.New("URL cannot be blank")
	}

	return Title(title), nil
}
