package main

import (
	"github.com/grafana/alertmanager-webhook-proxy/pkg/proxy"
	"github.com/grafana/alertmanager-webhook-proxy/pkg/server"
)

func main() {
	// TODO: Add configurable target host URL
	p, pErr := proxy.New("http://localhost:8090")
	srv := server.New(":8080", p)

	if pErr != nil {
		panic(pErr)
	}

	srvErr := server.Run(srv)

	if srvErr != nil {
		panic(srvErr)
	}
}
