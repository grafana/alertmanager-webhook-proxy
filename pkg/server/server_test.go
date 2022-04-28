package server

import (
	"testing"

	"github.com/grafana/alertmanager-webhook-proxy/pkg/proxy"
)

type test struct {
	addr string
}

var tests = []test{
	test{":1989"},
	test{":13"},
	test{":22"},
}

func TestNew(t *testing.T) {
	for _, test := range tests {
		p, _ := proxy.New("https://taylorswift.com")
		srv := New(test.addr, p)

		got := srv.Addr
		want := test.addr

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}
