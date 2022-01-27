package model

import (
	"bookshelf/domain/value"
)

type Memo struct {
	ID      int
	Content value.Content
}

// 新しいMemoモデルを作成します
func NewMemo(c string) (*Memo, error) {
	content, err := value.NewContent(c)
	if err != nil {
		return nil, err
	}
	return &Memo{ID: 0, Content: content}, nil
}

// メモの内容を編集します
func (m *Memo) SetContent(c string) error {
	replaced, err := value.NewContent(c)
	if err != nil {
		return err
	}
	m.Content = replaced
	return nil
}
