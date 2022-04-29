package main

import (
	"github.com/grafana/alertmanager-webhook-proxy/pkg/proxy"
	"github.com/grafana/alertmanager-webhook-proxy/pkg/server"
	"github.com/grafana/alertmanager-webhook-proxy/pkg/templater"
)

func main() {
	// TODO: Add configurable target host URL
	tmpl, _ := templater.New("/tmp/awp/template.txt")
	p, pErr := proxy.New("http://localhost:8090", tmpl)
	srv := server.New(":8080", p)

	if pErr != nil {
		panic(pErr)
	}

	srvErr := server.Run(srv)

	if srvErr != nil {
		panic(srvErr)
	}
}
