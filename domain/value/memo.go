package value

import (
	"errors"
	"unicode/utf8"
)

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
