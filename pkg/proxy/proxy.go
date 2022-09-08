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
	"strings"
	"text/template"

	"github.com/grafana/alertmanager-webhook-proxy/pkg/templater"
	amt "github.com/prometheus/alertmanager/template"
)

type ArrayFlag []string

func (i *ArrayFlag) String() string {
	return ""
}

func (i *ArrayFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func New(
	targetHost string,
	tmpl *template.Template,
	headers ArrayFlag,
) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)

	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	hMap := make(map[string]string)
	for _, h := range headers {
		s := strings.Split(h, ": ")
		hMap[s[0]] = s[1]
	}

	log.Println(hMap)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		modifyRequest(req, tmpl, hMap)
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

func modifyRequest(
	req *http.Request,
	tmpl *template.Template,
	hMap map[string]string,
) {
	d, err := Decode(req.Body)

	if err != nil {
		log.Println(err)
		return
	}

	s, _ := templater.Render(tmpl, d)

	req.ContentLength = int64(len(s))
	req.Body = ioutil.NopCloser(bytes.NewBufferString(s))

	for k, v := range hMap {
		log.Printf("Header: %v %v", k, v)
		req.Header.Set(k, v)
	}
}
