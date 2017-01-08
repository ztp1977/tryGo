package http

import (
	"bytes"
	"strings"
)

func Concat(ss ...string) string {
	var r string
	for _, s := range ss {
		r += s
	}

	return r
}

func Concat2(ss ...string) string {
	b := bytes.NewBufferString("")
	for _, s := range ss {
		r := strings.NewReader(s)
		r.WriteTo(b)
	}

	return b.String()
}
