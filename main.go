package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Bookmark struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

var id int
var inMemoryData []Bookmark

func main() {
	var bookmark http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf8")
		if r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			s, err := json.Marshal(inMemoryData)
			if err != nil {
				panic(err)
			}
			w.Write(s)
		}
		if r.Method == "POST" {
			var bookmarkPostBody struct {
				Url string `json:"url"`
			}
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			if err := json.Unmarshal(body, &bookmarkPostBody); err != nil {
				panic(err)
			}
			created := &Bookmark{id, bookmarkPostBody.Url}
			inMemoryData = append(inMemoryData, *created)
			w.Header().Set("Location", fmt.Sprintf("http://localhost:8080/bookmark/%d", created.Id))
			w.WriteHeader(http.StatusCreated)
			s, _ := json.Marshal(*created)
			w.Write(s)
			id++
		}
	}
	http.HandleFunc("/bookmark", bookmark)
	http.ListenAndServe(":8080", nil)
}
