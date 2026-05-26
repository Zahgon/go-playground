//go:build js
// +build js

package worker

import (
	"syscall/js"
)

type void = struct{}

// Releaser interface represents a value that should be released.
type Releaser interface {
	Release()
}

// Handlers is a key-value pair of function name and handler.
type Handlers = map[string]Func

// Worker is Go WebAssembly daemon that receives function calls
// from JS world and manages application lifecycle.
type Worker struct {
	done         chan void
	name         string
	exportObject map[string]any
	releaseList  []Releaser
}

// NewWorker constructs a new worker
func NewWorker() *Worker { _ = "STUB: not implemented"; return nil }

// RegisterFunc registers a function handler that will be exported.
func (w *Worker) RegisterFunc(fnName string, handler Func) { _ = "STUB: not implemented"; return }

// RegisterReleaser registers a releasable resource that should be released
// on worker shutdown by Worker.Release call.
//
// Used to register resources like js.Func that should be cleaned on shutdown.
func (w *Worker) RegisterReleaser(releaser Releaser, other ...Releaser) {
	_ = "STUB: not implemented"
	return
}

// Export calls a passed JS function with passed export object
// which contains all registered functions as argument.
//
// Exported object contains all functions added by Worker.RegisterFunc.
func (w *Worker) Export(callback js.Value) { _ = "STUB: not implemented"; return }

// Wait blocks the calling thread and waits for incoming calls
// until an exported exit function is called from JavaScript.
func (w Worker) Wait() {
	_ = "STUB: not implemented"

	// Release releases all worker resources.
	//
	// Should be called at worker shutdown.
	return
}

func (w Worker) Release() { _ = "STUB: not implemented"; return }

func (w Worker) onExit(this js.Value, args []js.Value) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// GetModuleExportCallback obtains a callback function to pass worker export object
// from worker command line arguments.
//
// Obtained function should be used as argument for Worker.Export method.
//
// Passed function name should be global.
func GetModuleExportCallback() (js.Value, error) {
	_ = "STUB: not implemented"
	return *new(js.Value), nil
}

// Exports is module exports map with key-value pair of name and function.
type Exports map[string]Func

// ExportAndStart starts a simple worker and exports functions it to JavaScript.
//
// For more sophisticated cases, use worker.NewWorker to manually
// create and register a worker.
//
// Also see worker.GetModuleExportCallback for more information
// about worker registration.
func ExportAndStart(exports Exports) { _ = "STUB: not implemented"; return }
