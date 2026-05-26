package index

import (
	"go/ast"
	"go/token"

	"github.com/x1unix/go-playground/internal/pkgindex/docutil"
)

var ignoreBuiltins = docutil.NewIgnoreList(
	"Type", "Type1", "IntegerType", "FloatType", "ComplexType",
)

type CollectFn = func(src SymbolSource, sym docutil.Symbol)

type sourceSummary struct {
	packageName  string
	symbolsCount int
	doc          *ast.CommentGroup
}

type fileParseParams struct {
	importPath string
	parseDoc   bool
	collector  CollectFn
}

func getFilter(importPath string) docutil.Filter {
	_ = "STUB: not implemented"
	return *new(docutil.Filter)
}

func parseFile(fset *token.FileSet, fpath string, params fileParseParams) (*sourceSummary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// "go/doc" ignores some packages from GOROOT thus it doesn't work for us.
// That means, all boring job should be done manually.
