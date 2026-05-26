package gowasm

import (
	"sync"

	"github.com/x1unix/go-playground/internal/util/syncx"
)

var (
	callbacks = syncx.NewMap[CallbackID, chan Result]()

	lastSeatID = 0
	lock       sync.Mutex
)

type (
	CallbackID = int
	Result     = int
)

// RequestCallback requests a new callback to await for async operation.
func RequestCallback() CallbackID { _ = "STUB: not implemented"; return *new(CallbackID) }

// AwaitCallback await for async operation to complete using callback ID.
func AwaitCallback(cbID CallbackID) error { _ = "STUB: not implemented"; return nil }

// AwaitResult awaits for callback completion and returns raw result.
func AwaitResult(cbID CallbackID) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

// ReleaseCallback releases acquired callback.
func ReleaseCallback(cbId CallbackID) { _ = "STUB: not implemented"; return }

// NotifyResult publishes callback result and notifies listener.
//
// This method is intended to be called only from JS side and kept public only for testing purposes.
func NotifyResult(cb CallbackID, result Result) { _ = "STUB: not implemented"; return }
