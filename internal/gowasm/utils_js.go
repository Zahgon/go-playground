package gowasm

import (
	"syscall/js"
)

// CopyBytesToGo is a panic-safe version of js.CopyBytesToGo
func CopyBytesToGo(dst []byte, src js.Value) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TODO: replace with custom panic-less implementation
