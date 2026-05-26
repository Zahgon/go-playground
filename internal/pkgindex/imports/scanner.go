// Package imports implements functionality for generating LSP completion data
// from documentation and symbols extracted from Go source files.
package imports

import (
	"regexp"

	"typefox.dev/lsp"
)

const (
	queueInitSize = 300

	// resultsInitSize value based on a number of packages in Go 1.22
	resultsInitSize = 180
)

var buildContraintRegex = regexp.MustCompile(`(?m)^([\S]+)_(freebsd|darwin|plan9|windows|netbsd|arm|386|loong64|mips|ppc|riscv|s390x)`)

type scanResult struct {
	hasPkg   bool
	item     lsp.CompletionItem
	children []string
}

// GoRootScanner scans Go SDK directory and provides information about Go version and standard packages list.
type GoRootScanner struct {
	goRoot string
}

func NewGoRootScanner(goRoot string) GoRootScanner {
	_ = "STUB: not implemented"
	return *new(GoRootScanner)
}

func (s *GoRootScanner) Scan() (*GoRootSummary, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *GoRootScanner) start() ([]lsp.CompletionItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *GoRootScanner) visitPackage(rootDir string, importPath string) (scanResult, error) {
	_ = "STUB: not implemented"
	return *new(scanResult), nil
}

func CheckVersionFile(root string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func shouldIgnoreFileName(fname string) bool { _ = "STUB: not implemented"; return false }

// isBuildConstraintFile checks whether file name contains Go build constraint.
//
// Linux and "_unix.go" files are allowed to align with GoDoc behavior.
// Wasm-target files are intentionally allowed to support `syscall/js` package.
//
// For example: `foo_amd64.go` or `bar_openbsd_ppc64.go`.
func isBuildConstraintFile(fname string) bool { _ = "STUB: not implemented"; return false }
