// Package check checks provided Go code and reports syntax errors
package check

import (
	"go/scanner"

	"typefox.dev/lsp"
)

func errorsListToMarkers(errList scanner.ErrorList) []lsp.Diagnostic {
	_ = "STUB: not implemented"
	return nil
}

// token.Position uses 1-based line/column indexes while LSP positions are 0-based.

func normalizeLSPPosition(value int) uint32 { _ = "STUB: not implemented"; return 0 }

// Result is result
type Result struct {
	// HasErrors is error status
	HasErrors bool `json:"hasErrors"`

	// Markers is list of diagnostics
	Markers []lsp.Diagnostic `json:"markers"`
}
