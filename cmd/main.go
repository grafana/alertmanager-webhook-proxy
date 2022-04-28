package main

import (
	"github.com/grafana/alertmanager-webhook-proxy/pkg/proxy"
)

func main() {
	// TODO: Add configurable target host URL
	p, err := proxy.New("http://localhost:8090")

	if err != nil {
		panic(err)
	}

	// TODO: Add configurable listening port
	proxy.Run(p, ":8080")
}
