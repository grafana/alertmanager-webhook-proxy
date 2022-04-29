package server

import (
	"net/http"
	"net/http/httputil"
)

func New(addr string, proxy *httputil.ReverseProxy) http.Server {
	srv := http.Server{
		Addr:    addr,
		Handler: proxy,
	}

	return srv
}

func Run(srv http.Server) error {
	err := srv.ListenAndServe()

	return err
}
