package proxy

import (
	"testing"
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
		_, err := New(test.target)

		if err != nil {
			t.Errorf("Error: %v", err)
		}
	}
}
