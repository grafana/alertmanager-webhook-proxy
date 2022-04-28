package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func New(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)

	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		modifyRequest(req)
	}

	return proxy, nil
}

func modifyRequest(req *http.Request) {
	// TODO: Add request body translator
}
