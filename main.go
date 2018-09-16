package main

import (
	"github.com/fucangyu/mblog/blog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", blog.Index)
	mux.HandleFunc("/upload/", blog.Auth(blog.Upload))
	mux.HandleFunc("/post/", blog.ReadPost)

	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: mux,
	}

	server.ListenAndServe()
}
