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

	server := blog.Server{Mux: mux}
	if blog.MODE == "HTTP" {
		server.ServeHTTP()
	} else if blog.MODE == "HTTPS" {
		server.ServeHttps()
	} else {
		panic("no http mode")
	}

}
