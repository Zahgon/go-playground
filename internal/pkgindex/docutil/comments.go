package docutil

import (
	"go/ast"
	"go/doc/comment"
)

const (
	pkgDocPrefix = "package "
	goDocBaseUrl = "https://pkg.go.dev/"

	// Char count of static markup chars from writeGoDocLink.
	goDocCharsLen = 22
)

// BuildPackageDoc builds markdown documentation for package with a link to GoDoc page.
//
// Please call [IsPackageDoc] before using this method.
func BuildPackageDoc(group *ast.CommentGroup, importPath string) string {
	_ = "STUB: not implemented"
	return ""
}

// EmptyPackageDoc returns empty package doc which only contains link to GoDoc.
func EmptyPackageDoc(importPath string) string { _ = "STUB: not implemented"; return "" }

func appendGoDocLink(docStr []byte, importPath string) string {
	_ = "STUB: not implemented"
	// Approx new buffer length. 20 is length of static characters.
	return ""
}

// [pkgName on pkg.go.dev](https://pkg.go.dev/importPath)

// FormatCommentGroup parses comments from AST and returns them in Markdown format.
func FormatCommentGroup(group *ast.CommentGroup) []byte { _ = "STUB: not implemented"; return nil }

func replaceDocLinks(doc *comment.Doc) map[string]string { _ = "STUB: not implemented"; return nil }

func replaceDocLinksInBlock(block comment.Block, replacements map[string]string, nextID *int) comment.Block {
	_ = "STUB: not implemented"
	return *new(comment.Block)
}

func replaceDocLinksInText(text []comment.Text, replacements map[string]string, nextID *int) []comment.Text {
	_ = "STUB: not implemented"
	return nil
}

func extractText(text []comment.Text) string { _ = "STUB: not implemented"; return "" }

func replacePlainSymbolRefs(text string, replacements map[string]string, nextID *int) []comment.Text {
	_ = "STUB: not implemented"
	return nil
}

func isBoundaryBefore(text string, i int) bool { _ = "STUB: not implemented"; return false }

func isBoundaryAfter(text string, i int) bool { _ = "STUB: not implemented"; return false }

func isSymbolRef(ref string) bool { _ = "STUB: not implemented"; return false }

func applyReplacements(mdDoc []byte, replacements map[string]string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// IsPackageDoc returns whether top level comment is a valid package comment.
//
// See issue #367.
func IsPackageDoc(group *ast.CommentGroup) bool { _ = "STUB: not implemented"; return false }

// We interested only in first non-empty comment.

func hasPackagePrefix(c string) bool { _ = "STUB: not implemented"; return false }

// isDirective reports whether c is a comment directive.
//
// Copy from go/ast/ast.go:172
func isDirective(c string) bool {
	_ = "STUB: not implemented"
	// "//line " is a line directive.
	// "//extern " is for gccgo.
	// "//export " is for cgo.
	// (The // has been removed.)
	return false
}

// "//[a-z0-9]+:[a-z0-9]"
// (The // has been removed.)
