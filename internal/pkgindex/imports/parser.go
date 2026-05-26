package imports

import (
	"context"

	"typefox.dev/lsp"
)

type PackageParseParams struct {
	RootDir    string
	ImportPath string
	Files      []string
}

func (params PackageParseParams) PackagePath() string { _ = "STUB: not implemented"; return "" }

// ParseImportCompletionItem parses a Go package at a given GOROOT and constructs an LSP completion item from it.
func ParseImportCompletionItem(ctx context.Context, params PackageParseParams) (result lsp.CompletionItem, error error) {
	_ = "STUB: not implemented"
	return *new(lsp.CompletionItem), nil
}

// Package doc should start with "Package xxx", see issue #367.
