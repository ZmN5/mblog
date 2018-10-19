package blog

import (
	"crypto/tls"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
)

type Server struct {
	Mux *http.ServeMux
}

func (server Server) ServeHttps() {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(DOMAIN),
		Cache:      autocert.DirCache(CERTS),
	}

	httpsServer := &http.Server{
		Addr:    ":https",
		Handler: server.Mux,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}
	go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

	log.Fatal(httpsServer.ListenAndServeTLS("", ""))

}

func (server Server) ServeHTTP() {

	httpServer := &http.Server{
		Addr:    ":http",
		Handler: server.Mux,
	}

	log.Fatal(httpServer.ListenAndServe())
}
