package proxy

import (
	"bytes"
	"io/ioutil"
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

func TestDecode(t *testing.T) {
	payload, rErr := ioutil.ReadFile("testdata/sample.json")

	if rErr != nil {
		t.Error("failed to read sample payload")
	}

	reader := bytes.NewReader(payload)

	data, dErr := decode(reader)

	if dErr != nil {
		t.Errorf("failed to decode: %v", dErr)
	}

	got := data.Receiver
	want := "webhook"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
