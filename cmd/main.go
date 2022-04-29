package main

import (
	"flag"

	"github.com/grafana/alertmanager-webhook-proxy/pkg/proxy"
	"github.com/grafana/alertmanager-webhook-proxy/pkg/server"
	"github.com/grafana/alertmanager-webhook-proxy/pkg/templater"
)

func main() {
	tmplPtr := flag.String("template", "/tmp/awp/template.txt", "Path to alert payload transformation template")
	targetPtr := flag.String("target", "http://localhost:8090", "Target URL")
	addressPtr := flag.String("address", ":8080", "Server bind address")

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
