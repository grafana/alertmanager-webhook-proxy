package proxy

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/prometheus/alertmanager/template"
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

func decode(body io.Reader) (template.Data, error) {
	if body == nil {
		return template.Data{}, nil
	}

	decoder := json.NewDecoder(body)

	var data template.Data

	err := decoder.Decode(&data)

	return data, err
}

func modifyRequest(req *http.Request) {
	_, err := decode(req.Body)

	if err != nil {
		log.Println(err)
		return
	}
}
