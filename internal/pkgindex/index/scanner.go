package index

import (
	"github.com/x1unix/go-playground/internal/pkgindex/imports"
)

const (
	// queueSize is based on max occupation of a queue during test scan of Go 1.23.
	//
	// See: Queue.MaxOccupancy
	queueSize = 120

	// Go 1.23 has 182 packages and over 9k total symbols for linux.
	pkgBuffSize = 182
	symBuffSize = 9000
)

var Debug = false

type scanEntry struct {
	isVendor   bool
	path       string
	importPath string
}

func ScanRoot(goRoot string) (*GoIndexFile, error) { _ = "STUB: not implemented"; return nil, nil }

// populate queue with root packages

// Edge case: Apparently GOROOT has vendoring for its own packages.

// Edge case: "builtin" package exists only for documentation purposes
// and not importable.
// Also skip empty packages (usually part of vendor path).

func enqueueRootEntries(rootDir string, parentImportPath string, queue *imports.Queue[scanEntry]) error {
	_ = "STUB: not implemented"
	return nil
}

func isDirIgnored(basename string) bool { _ = "STUB: not implemented"; return false }

func isImportPathIgnored(importPath string) bool {
	_ = "STUB: not implemented"
	// Arena experiment was rejected and removed
	return false
}
