package docutil

import (
	"go/ast"
	"go/token"
	"reflect"
	"strings"

	"typefox.dev/lsp"
)

const defaultStringBuffSize = 64

const BuiltinPackage = "builtin"

var (
	astDocFields     = []string{"Doc", "Comment"}
	commentBlockType = reflect.TypeFor[*ast.CommentGroup]()

	token2KindMapping = map[token.Token]lsp.CompletionItemKind{
		token.VAR:   lsp.VariableCompletion,
		token.CONST: lsp.ConstantCompletion,
		token.TYPE:  lsp.ClassCompletion,
	}
)

// TokenToCompletionItemKind maps Go AST token to LSP completion item kind.
func TokenToCompletionItemKind(tok token.Token) (lsp.CompletionItemKind, bool) {
	_ = "STUB: not implemented"
	return *new(lsp.CompletionItemKind), false
}

// IsGoSourceFile returns whether a file name is a Go source file.
//
// Returns false for unit test files (this is intentional behavior).
func IsGoSourceFile(name string) bool { _ = "STUB: not implemented"; return false }

// CommentGroupEmpty checks whether passed command group is empty.
func CommentGroupEmpty(g *ast.CommentGroup) bool { _ = "STUB: not implemented"; return false }

// GetDeclRange returns AST node range in document.
func GetDeclRange(fset *token.FileSet, decl ast.Decl) (start token.Position, end token.Position) {
	_ = "STUB: not implemented"
	return *new(token.Position), *new(token.Position)
}

type Positioner interface {
	Pos() token.Pos
}

// GetDeclPosition returns start position of an AST node
func GetDeclPosition(fset *token.FileSet, decl Positioner) token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

// OffsetToPosition translates offset into position with column and line number.
func OffsetToPosition(fset *token.FileSet, pos token.Pos) token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

// WriteTypeParams writes type parameters snippet template into a specified buffer.
//
// Snippet offset is a start index for snippet template variables (`$n`) to fill parameters.
func WriteTypeParams(sb *strings.Builder, snippetIndex int, typeParams *ast.FieldList) int {
	_ = "STUB: not implemented"
	return 0
}

// WriteParamsList writes parameters list template (usually func args) into a specified buffer.
//
// Snippet offset is a start index for snippet template variables (`$n`) to fill parameters.
func WriteParamsList(sb *strings.Builder, snippetIndex int, params *ast.FieldList) int {
	_ = "STUB: not implemented"
	return 0
}

// PrintDecl returns string representation of passed AST node.
func PrintDecl(fset *token.FileSet, decl any) (string, error) {
	_ = "STUB: not implemented"
	// Remove comments block from AST node to keep only node body
	return "", nil
}

func removeCommentFromDecl(decl any) any { _ = "STUB: not implemented"; return *new(any) }

// *ast.FuncDecl, *ast.Object have Doc
// *ast.Object and *ast.Indent might have Comment
