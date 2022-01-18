package main

import (
	"encoding/json"
	"net/http"
)

type Bookmark struct {
	Url string `json:"url"`
}

func main() {
	sampleBookmark := Bookmark{"https://youtube.com"}
	http.HandleFunc("/bookmark", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST")
		if r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			s, err := json.Marshal(sampleBookmark)
			if err != nil {
				panic(err)
			}
			w.Write(s)
		}
	})

	http.ListenAndServe(":8080", nil)
}
