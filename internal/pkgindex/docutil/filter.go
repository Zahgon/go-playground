package docutil

import (
	"github.com/hashicorp/go-set/v3"
)

// Filter is interface to provide way to ignore certain symbols during AST traversal.
type Filter interface {
	Ignore(typeName string) bool
}

type ignoreList struct {
	m *set.Set[string]
}

func (f ignoreList) Ignore(typeName string) bool { _ = "STUB: not implemented"; return false }

// NewIgnoreList creates a filter with a list of ignored symbols.
func NewIgnoreList(names ...string) Filter { _ = "STUB: not implemented"; return *new(Filter) }

// UnexportedFilter filters private symbols
type UnexportedFilter struct{}

func (UnexportedFilter) Ignore(typeName string) bool { _ = "STUB: not implemented"; return false }

type composedFilter []Filter

func (filters composedFilter) Ignore(typeName string) bool { _ = "STUB: not implemented"; return false }

// ComposeFilters allows composing multiple symbol filters into one.
func ComposeFilters(filters ...Filter) Filter { _ = "STUB: not implemented"; return *new(Filter) }

func filterOrDefault(f Filter) Filter { _ = "STUB: not implemented"; return *new(Filter) }
