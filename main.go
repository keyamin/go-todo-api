package main

import (
	"fmt"
	"net/http"
)

var bookmarks []Bookmark

type Bookmark struct {
	Url string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "sample")
	})

	http.ListenAndServe(":8080", nil)
}

func PostBookmark() {
	var url string
	fmt.Scan(&url)
	bookmarks = append(bookmarks, Bookmark{url})
}

func DeleteBookmark() {
	fmt.Println("Select index you want to delete...")
	for index, value := range bookmarks {
		fmt.Printf("%d: %s\n", index+1, value.Url)
	}
	var index int
	fmt.Scan(&index)
	index -= 1
	bookmarks = bookmarks[:index+copy(bookmarks[index:], bookmarks[index+1:])]
	fmt.Println("deleted!!")
}
