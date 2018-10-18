package main

import (
	"crypto/tls"
	"github.com/fucangyu/mblog/blog"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
)

func main() {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(blog.DOMAIN), //Your domain here
		Cache:      autocert.DirCache(blog.CERTS),       //Folder for storing certificates
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", blog.Index)
	mux.HandleFunc("/upload/", blog.Auth(blog.Upload))
	mux.HandleFunc("/post/", blog.ReadPost)

	server := &http.Server{
		Addr:    ":https",
		Handler: mux,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}
	go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

	log.Fatal(server.ListenAndServeTLS("", ""))
}
