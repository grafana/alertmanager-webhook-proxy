package main

import (
	"flag"

	"github.com/grafana/alertmanager-webhook-proxy/pkg/proxy"
	"github.com/grafana/alertmanager-webhook-proxy/pkg/server"
	"github.com/grafana/alertmanager-webhook-proxy/pkg/templater"
)

func main() {
	addressPtr := flag.String("address", ":8080", "Server bind address")
	targetPtr := flag.String("target", "http://localhost:8090", "Target URL")
	tmplPtr := flag.String("template", "/tmp/awp/template.txt", "Path to payload transformation template")

	flag.Parse()

	tmpl, _ := templater.New(*tmplPtr)
	p, pErr := proxy.New(*targetPtr, tmpl)
	srv := server.New(*addressPtr, p)

	if pErr != nil {
		panic(pErr)
	}

	srvErr := server.Run(srv)

	if srvErr != nil {
		panic(srvErr)
	}
}
