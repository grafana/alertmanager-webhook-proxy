package templater

import (
	"bytes"
	"text/template"

	amt "github.com/prometheus/alertmanager/template"
)

func New(path string) (*template.Template, error) {
	t, err := template.ParseFiles(path)

	return t, err
}

func Render(t *template.Template, data amt.Data) (string, error) {
	var b bytes.Buffer

	err := t.Execute(&b, data)

	return b.String(), err
}
