package main

import (
	"bytes"
	"path/filepath"
	"text/template"
)

func getWithTemplate(data interface{}, files ...string) ([]byte, error) {
	f := make([]string, len(files))
	for i, p := range files {
		f[i] = filepath.Join("templates", p)
	}

	tmpl := template.Must(template.ParseFiles(f...))

	buf := &bytes.Buffer{}
	err := tmpl.Execute(buf, data)

	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}
