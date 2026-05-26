package docutil

import (
	"go/ast"
	"go/token"
)

type TraverseContext struct {
	FileSet *token.FileSet
	Block   BlockData
	Filter  Filter
}

// CollectValues constructs completion item from value declaration.
//
// Able to handle special edge cases for builtin declarations.
func CollectValues(ctx TraverseContext, spec *ast.ValueSpec, collector Collector) (count int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getValueDocumentation(block BlockData, spec *ast.ValueSpec) string {
	_ = "STUB: not implemented"
	return ""
}

func signatureFromIdent(fset *token.FileSet, block BlockData, ident *ast.Ident) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// isIdentPrintable checks whether completion item detail can be
// generated from identifier.
//
// Returns stripped, save to print version of identifier value on success.
//
// Returns true if identifier has type definition, scalar value or is from "builtin" package.
func isIdentPrintable(ident *ast.Ident) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Edge case: builtin package.

// Permit only value declarations

// Allow identifier with type

// Allow typed identifiers, but strip value
// Return type should be a pointer in order to satisfy printer interfaces.
