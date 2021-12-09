package main

import (
	"fmt"
	"io/ioutil"
	"testing"
	"text/template"
)

func BenchmarkTemplate(b *testing.B) {
	format := "[{{.Context}}/{{.Namespace}}]"

	for i := 0; i < b.N; i++ {
		// The template compilation is run on any execution of this binary,
		// so include this in the benchmark
		template.Must(template.New("tmux").Parse(format)).Execute(ioutil.Discard, struct {
			Context, Namespace string
		}{
			"dummy-ctx",
			"dummy-ns",
		})
	}
}

func BenchmarkPrinting(b *testing.B) {
	format := "[%s/%s]"

	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, format, "dummy-ctx", "dummy-ns")
	}
}
