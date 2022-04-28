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

func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func modifyRequest(req *http.Request) {
	// TODO: Add request body translator
}

func Run(proxy *httputil.ReverseProxy, addr string) {
	http.HandleFunc("/", ProxyRequestHandler(proxy))

	http.ListenAndServe(addr, nil)
}
