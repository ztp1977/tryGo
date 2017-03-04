package main

import (
	"html/template"

	"github.com/k0kubun/pp"
)

func main() {
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
	pp.Println(err)

	pp.Println(out)
}
