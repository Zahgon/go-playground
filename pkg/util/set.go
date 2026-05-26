package util

type void = struct{}

// StringSet is a list of unique strings.
type StringSet map[string]void

// Has checks if set contains a value.
func (s StringSet) Has(val string) bool { _ = "STUB: not implemented"; return false }

// Concat joins two string sets together into a new set.
func (s StringSet) Concat(items StringSet) StringSet {
	_ = "STUB: not implemented"
	return *new(StringSet)
}

// NewStringSet creates a new string set from strings slice.
func NewStringSet(items ...string) StringSet { _ = "STUB: not implemented"; return *new(StringSet) }
