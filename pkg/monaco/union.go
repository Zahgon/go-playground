package monaco

import (
	"encoding/json"
)

var (
	_ json.Marshaler   = (*UnionString[struct{}])(nil)
	_ json.Unmarshaler = (*UnionString[struct{}])(nil)
)

// UnionString is a special union type to represent values that can be a string or something else.
//
// For example:
//
//	interface CompletionItem {
//		label: string | CompletionItemLabel;
//	}
//
// Can be represented in Go as:
//
//	type CompletionItem struct {
//		Label UnionString[CompletionItemLabel] `json:"label,omitempty"`
//	}
//
// To initialize a union, fill a proper struct or string fields:
//
//	strValue := monaco.UnionString[monaco.CompletionItemLabel]{
//		String: "raw string value",
//	}
//
//	objValue := monaco.UnionString[monaco.CompletionItemLabel]{
//		Value: &monaco.CompletionItemLabel{
//			Label: "a label",
//		},
//	}
//
// Helper methods can be used to quickly replace values in a union.
//
//	var item monaco.CompletionItem
//	item.Label.SetString("raw string")
//	item.Label.SetValue(&CompletionItemLabel{...})
type UnionString[T any] struct {
	// String is a raw string value in a union.
	String string

	// Value is alternative union value.
	Value *T
}

// SetString replaces union value with a raw string.
func (u *UnionString[T]) SetString(str string) { _ = "STUB: not implemented"; return }

// SetValue replaces union value with a object.
func (u *UnionString[T]) SetValue(v *T) { _ = "STUB: not implemented"; return }

// MarshalJSON implements json.Marshaler interface.
func (u UnionString[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler interface.
func (u *UnionString[T]) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// null or undefined
