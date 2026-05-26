package gowasm

import (
	"context"
	"syscall/js"
)

type Func func(ctx context.Context, args []js.Value) (any, error)

type Worker struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	exports    map[string]Func
}

// NewWorker constructs a new worker object.
func NewWorker() *Worker { _ = "STUB: not implemented"; return nil }

// Context returns worker execution context
func (w *Worker) Context() context.Context {
	_ = "STUB: not implemented"

	// Export adds a new callback to worker export object
	return *new(context.Context)
}

func (w *Worker) Export(methodName string, cb Func) { _ = "STUB: not implemented"; return }

type jsFunc struct {
	ref   uint64
	grPtr uintptr
	id    uint32
}

// Run registers worker and starts program execution.
//
// This method blocks the main goroutine until worker context is alive.
func (w *Worker) Run() { _ = "STUB: not implemented"; return }

// handleCall handles calls from JS and routes them to handlers.
func (w *Worker) handleCall(_ js.Value, args []js.Value) any {
	_ = "STUB: not implemented"
	return *new(any)
}
