//go:build js
// +build js

// Package worker contains Go web-worker WASM module bridge methods
package worker

import (
	"syscall/js"
)

// Func is worker handler function
type Func = func(this js.Value, args Args) (interface{}, error)

// ParseArgs parses async call arguments.
//
// Function expects the last argument to be a callable JS function
func ParseArgs(allArgs []js.Value) (Args, Callback, error) {
	_ = "STUB: not implemented"
	return *new(Args), *new(Callback), nil
}

func callFunc(fn Func, this js.Value, jsArgs []js.Value) { _ = "STUB: not implemented"; return }

// FuncOf wraps function into js-compatible async function with callback
func FuncOf(fn Func) js.Func { _ = "STUB: not implemented"; return *new(js.Func) }
