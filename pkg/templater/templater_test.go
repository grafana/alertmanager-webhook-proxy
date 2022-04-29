package templater

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	amt "github.com/prometheus/alertmanager/template"
)

func TestMain(t *testing.M) {
	v := t.Run()

	snaps.Clean()

	os.Exit(v)
}

func TestRender(t *testing.T) {
	tmpl, err := New("testdata/template.txt")

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	payload, rErr := ioutil.ReadFile("testdata/sample.json")

	if rErr != nil {
		t.Error("failed to read sample payload")
	}

	var data amt.Data
	reader := bytes.NewReader(payload)
	decoder := json.NewDecoder(reader)
	dErr := decoder.Decode(&data)

	if dErr != nil {
		t.Errorf("failed to decode: %v", dErr)
	}

	s, reErr := Render(tmpl, data)

	if reErr != nil {
		t.Errorf("failed to render: %v", reErr)
	}

	if !json.Valid([]byte(s)) {
		t.Errorf("invalid json: %v", s)
	}

	snaps.MatchSnapshot(t, s)
}
