package proxy

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"text/template"

	"github.com/grafana/alertmanager-webhook-proxy/pkg/templater"
	amt "github.com/prometheus/alertmanager/template"
)

func New(targetHost string, tmpl *template.Template) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)

	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		modifyRequest(req, tmpl)
	}

	return proxy, nil
}

func Decode(body io.Reader) (amt.Data, error) {
	if body == nil {
		return amt.Data{}, nil
	}

	decoder := json.NewDecoder(body)

	var data amt.Data

	err := decoder.Decode(&data)

	return data, err
}

func modifyRequest(req *http.Request, tmpl *template.Template) {
	d, err := Decode(req.Body)

	if err != nil {
		log.Println(err)
		return
	}

	s, _ := templater.Render(tmpl, d)

	req.ContentLength = int64(len(s))
	req.Body = ioutil.NopCloser(bytes.NewBufferString(s))
}
