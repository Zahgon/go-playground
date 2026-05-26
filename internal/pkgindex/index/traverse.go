package index

import (
	"go/build"

	"github.com/x1unix/go-playground/internal/pkgindex/imports"
)

// defaultCtx is build constraint context to select only files
// that match desired platform.
//
// To keep aligned with godoc - uses same GOOS and GOARCH.
var defaultCtx = build.Context{
	GOARCH: "amd64",
	GOOS:   "linux",
}

type traverseResult struct {
	pkgInfo      PackageInfo
	symbolsCount int
}

func isFileIgnored(entry scanEntry, fname string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// provide docs for single-platform packages.

func traverseScanEntry(entry scanEntry, queue *imports.Queue[scanEntry], collector CollectFn) (*traverseResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func enqueueSubDir(queue *imports.Queue[scanEntry], parent scanEntry, name string) {
	_ = "STUB: not implemented"
	return
}

// TODO: should nested vendors be supported?
