package docutil

import (
	"go/ast"
	"go/token"

	"typefox.dev/lsp"
)

// SymbolFromFunc constructs completion item from a function AST declaration.
//
// Function documentation is generated in Markdown format.
func SymbolFromFunc(fset *token.FileSet, fn *ast.FuncDecl, snippetFormat lsp.InsertTextFormat) (item Symbol, err error) {
	_ = "STUB: not implemented"
	return *new(Symbol), nil
}

func buildFuncInsertStatement(decl *ast.FuncDecl, asSnippet bool) string {
	_ = "STUB: not implemented"
	return ""
}

// snippet offsets start at 1

// PrintFuncPrototype returns function string representation without its body.
func PrintFuncPrototype(fset *token.FileSet, decl *ast.FuncDecl) (string, error) {
	_ = "STUB: not implemented"
	// drop body from func
	return "", nil
}

// PrintFuncAnonymous similar to PrintFuncPrototype, but omits function name.
func PrintFuncAnonymous(fset *token.FileSet, decl *ast.FuncDecl) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
