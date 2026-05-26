//go:build js
// +build js

package worker

import (
	"syscall/js"
)

// NewTypeError creates a new type error
func NewTypeError(expType, gotType js.Type) error { _ = "STUB: not implemented"; return nil }

// ValueUnmarshaler unmarshal JS value
type ValueUnmarshaler interface {
	// UnmarshalValue unmarshal JS value
	UnmarshalValue(js.Value) error
}

// Args is collection if function call arguments
type Args []js.Value

// BindIndex binds argument at specified index to passed value
func (args Args) BindIndex(index int, dest interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Bind binds passed JS arguments to Go values
//
// Function supports *int, *bool, *string and ValueUnmarshaler values.
func (args Args) Bind(targets ...interface{}) error { _ = "STUB: not implemented"; return nil }

// BindValue binds JS value to specified target
func BindValue(val js.Value, dest interface{}) error { _ = "STUB: not implemented"; return nil }
