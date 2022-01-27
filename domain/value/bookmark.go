package value

import (
	"errors"
)

type Url string

func NewUrl(url string) (Url, error) {
	if url == "" {
		return "", errors.New("URL cannot be blank")
	}

	return Url(url), nil
}
