package main

import (
	"github.com/fucangyu/mblog/blog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", blog.Index)
	mux.HandleFunc("/upload/", blog.Upload)
	mux.HandleFunc("/post/", blog.ReadPost)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
