package docutil

import (
	"go/ast"
	"go/token"

	"typefox.dev/lsp"
)

// BlockData contains information about declaration group.
type BlockData struct {
	IsGroup bool
	Decl    *ast.GenDecl
	Kind    lsp.CompletionItemKind
}

// NewBlockData parses block data from AST block declaration node.
//
// Example:
//
//	var (
//		foo = 1
//		bar = "bar"
//	)
//
//	type (
//		foo int
//		bar struct{}
//	)
func NewBlockData(specGroup *ast.GenDecl) (BlockData, error) {
	_ = "STUB: not implemented"
	return *new(BlockData), nil
}

// TypeToSymbol returns completion item from type declaration inside block.
func TypeToSymbol(fset *token.FileSet, block BlockData, spec *ast.TypeSpec) (Symbol, error) {
	_ = "STUB: not implemented"
	// Block declarations contain doc inside each child.
	return *new(Symbol), nil
}

// TODO: prefill struct members

func getTypeDoc(block BlockData, spec *ast.TypeSpec) string { _ = "STUB: not implemented"; return "" }

// standalone type declarations are still considered as block.
