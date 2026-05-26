//go:build js
// +build js

package main

import (
	"syscall/js"

	"github.com/x1unix/go-playground/pkg/worker"
)

func main() {
	worker.ExportAndStart(worker.Exports{
		"analyzeCode": analyzeCode,
	})
}

func analyzeCode(this js.Value, args worker.Args) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
