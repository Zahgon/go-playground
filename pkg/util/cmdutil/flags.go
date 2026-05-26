package cmdutil

import (
	"encoding"
)

const csvSeparator = ","

// StringsListValue is comma-separated list of values that implements flag.Value interface.
type StringsListValue []string

// String implements flag.Value
func (s StringsListValue) String() string { _ = "STUB: not implemented"; return "" }

// Set implements flag.Value
func (s *StringsListValue) Set(s2 string) error { _ = "STUB: not implemented"; return nil }

// NewStringsListValue returns a new StringsListValue
func NewStringsListValue(p *[]string) *StringsListValue { _ = "STUB: not implemented"; return nil }

// TextUnmarshalerValue is flag.Value adapter for values
// which implement encoding.TextUnmarshaler interface.
type TextUnmarshalerValue struct {
	dest encoding.TextUnmarshaler
}

func NewTextUnmarshalerValue(dest encoding.TextUnmarshaler) *TextUnmarshalerValue {
	_ = "STUB: not implemented"
	return nil
}

func (t *TextUnmarshalerValue) WithDefaults(defaultValue string) *TextUnmarshalerValue {
	_ = "STUB: not implemented"
	return nil
}

// Set implements flag.Value
func (t TextUnmarshalerValue) Set(v string) error { _ = "STUB: not implemented"; return nil }

// String implements flag.Value
func (t TextUnmarshalerValue) String() string { _ = "STUB: not implemented"; return "" }
