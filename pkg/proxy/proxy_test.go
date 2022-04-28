package proxy

import (
	"testing"

	"net/http"
)

type test struct {
	target string
}

var tests = []test{
	test{"https://taylorswift.com"},
	test{"https://grafana.com"},
	test{"https://play.grafana.com"},
}

func TestNew(t *testing.T) {
	for _, test := range tests {
		p, err := New(test.target)

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		req, _ := http.NewRequest("POST", "https://replace.me", nil)

		p.Director(req)

		got := req.URL.Scheme + "://" + req.URL.Host + req.URL.Path
		want := test.target + "/"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}
