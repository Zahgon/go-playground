package docutil

import (
	"go/ast"
	"go/token"

	"typefox.dev/lsp"
)

type TraverseOpts struct {
	Filter        Filter
	FileSet       *token.FileSet
	SnippetFormat lsp.InsertTextFormat
}

// CollectSymbols traverses root file declarations and transforms them into completion items.
//
// Important: type methods are ignored.
func CollectSymbols(decls []ast.Decl, opts TraverseOpts, collector Collector) (count int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Ignore type methods, at-least for now.
