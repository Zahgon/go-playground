//go:build js
// +build js

package worker

import "syscall/js"

// Callback is async function callback
type Callback = func(interface{}, error)

func newCallbackFromValue(val js.Value) (Callback, error) {
	_ = "STUB: not implemented"
	return *new(Callback), nil
}
