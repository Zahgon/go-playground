package docutil

import (
	"go/ast"
	"go/token"
)

// CollectDecls collects symbols from generic type of value spec.
func CollectDecls(fset *token.FileSet, specGroup *ast.GenDecl, filter Filter, collector Collector) (count int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
