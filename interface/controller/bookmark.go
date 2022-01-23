package controller

import (
	"bookshelf/usecase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BookmarkController interface {
	Post() http.HandlerFunc
}

type BookmarkControllerImpl struct {
	bookmarkUsecase usecase.BookmarkUsecase
}

func NewBookmarkController(bookmarkUsecase usecase.BookmarkUsecase) BookmarkController {
	return &BookmarkControllerImpl{bookmarkUsecase}
}

type requestBookmark struct {
	Url string `json:"url"`
}
type responseBookmark struct {
	ID  int    `json:"id"`
	Url string `json:"url"`
}

func (b *BookmarkControllerImpl) Post() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf8")
		var bookmarkPostBody requestBookmark
		body, readErr := ioutil.ReadAll(r.Body)
		if readErr != nil {
			panic(readErr)
		}
		if err := json.Unmarshal(body, &bookmarkPostBody); err != nil {
			panic(err)
		}
		createdModel, usecaseErr := b.bookmarkUsecase.Create(bookmarkPostBody.Url)
		if usecaseErr != nil {
			panic(usecaseErr)
		}
		created := &responseBookmark{
			ID:  createdModel.ID,
			Url: createdModel.Url.GetUrl(),
		}
		w.Header().Set("Location", fmt.Sprintf("http://localhost:8080/bookmark/%d", created.ID))
		w.WriteHeader(http.StatusCreated)
		s, _ := json.Marshal(*created)
		w.Write(s)
	}
}
