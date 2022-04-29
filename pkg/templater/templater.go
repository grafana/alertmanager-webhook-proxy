package templater

import (
	"io"
	"text/template"

	amt "github.com/prometheus/alertmanager/template"
)

func New(path string) (*template.Template, error) {
	t, err := template.ParseFiles(path)

	return t, err
}

func Render(wr io.Writer, t *template.Template, data amt.Data) error {
	err := t.Execute(wr, data)

	return err
}
