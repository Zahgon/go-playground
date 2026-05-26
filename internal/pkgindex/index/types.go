package index

import (
	"github.com/x1unix/go-playground/internal/pkgindex/docutil"
	"typefox.dev/lsp"
)

const GoIndexFileVersion = 2

type FlatSymbolSource [2]string

type PackageInfo struct {
	// Name is package name.
	Name string `json:"name"`

	// ImportPath is full import path of a package.
	ImportPath string `json:"importPath"`

	// Doc is documentation in Markdown format.
	Doc string `json:"doc,omitempty"`
}

// Packages is a flat representation of PackageInfo list.
type Packages struct {
	Names []string `json:"names"`
	Paths []string `json:"paths"`
	Docs  []string `json:"docs"`
}

func NewPackages(capacity int) Packages { _ = "STUB: not implemented"; return *new(Packages) }

func (pkgs *Packages) Append(pkg PackageInfo) { _ = "STUB: not implemented"; return }

// Symbols is a flat representation of Go package symbols.
type Symbols struct {
	// Names are list of symbol names.
	Names []string `json:"names"`

	// Docs are list of symbol documentation.
	Docs []string `json:"docs,omitempty"`

	// Details are list of short symbol summaries
	Details []string `json:"details,omitempty"`

	// Signatures contain string representation of a symbol (e.g. struct definition)
	// which is visible when user hovers on a symbol.
	Signatures []string `json:"signatures,omitempty"`

	// InsertTexts are values to be inserted when symbol suggestion is selected.
	InsertTexts []string `json:"insertTexts"`

	// InsertTextRules contains snippet insertion rules for InsertTexts.
	InsertTextRules []lsp.InsertTextFormat `json:"insertTextRules,omitempty"`

	// Kinds contains symbol type for suggestion icon.
	Kinds []lsp.CompletionItemKind `json:"kinds"`

	// Packages contains information where particular symbol belongs (to what package).
	Packages []FlatSymbolSource `json:"packages"`
}

func NewSymbols(capacity int) Symbols { _ = "STUB: not implemented"; return *new(Symbols) }

func (s *Symbols) Append(src SymbolSource, sym docutil.Symbol) { _ = "STUB: not implemented"; return }

// SymbolSource holds information where symbol belongs to.
type SymbolSource struct {
	// Name is package name.
	Name string `json:"name"`

	// Path is import path of a package.
	Path string `json:"path"`
}

func (s SymbolSource) Flatten() FlatSymbolSource {
	_ = "STUB: not implemented"
	return *new(FlatSymbolSource)
}

// GoIndexFile contains flat list of all Go packages and symbols (functions, types and values).
//
// Data is organized into a flat structure (soa) in order to reduce output file size.
type GoIndexFile struct {
	// Version is file format version.
	Version int `json:"version"`

	// Go is Go version used to generate index.
	Go string `json:"go"`

	// Packages is structure of arrays of standard Go packages.
	Packages Packages `json:"packages"`

	// Symbols is structure of arrays of package symbols.
	Symbols Symbols `json:"symbols"`
}
