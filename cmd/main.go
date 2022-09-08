package main

import (
	"flag"
	"log"

	"github.com/grafana/alertmanager-webhook-proxy/pkg/proxy"
	"github.com/grafana/alertmanager-webhook-proxy/pkg/server"
	"github.com/grafana/alertmanager-webhook-proxy/pkg/templater"
)

func main() {
	var headers proxy.ArrayFlag

	addressPtr := flag.String("address", ":8080", "Server bind address")
	targetPtr := flag.String("target", "http://localhost:8090", "Target URL")
	tmplPtr := flag.String("template", "/tmp/awp/template.txt", "Path to payload transformation template")
	flag.Var(&headers, "header", "Header to set for the proxied request")

	flag.Parse()

	log.Printf("Bind address: %v", *addressPtr)
	log.Printf("Target address: %v", *targetPtr)
	log.Printf("Template path: %v", *tmplPtr)

	for _, h := range headers {
		log.Printf("Header: %v", h)
	}

	tmpl, _ := templater.New(*tmplPtr)
	p, pErr := proxy.New(*targetPtr, tmpl, headers)
	srv := server.New(*addressPtr, p)

	if pErr != nil {
		panic(pErr)
	}

	srvErr := server.Run(srv)

	if srvErr != nil {
		panic(srvErr)
	}
}
