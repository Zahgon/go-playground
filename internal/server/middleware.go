package server

import (
	"net/http"
	"time"
)

// HandlerFunc is langserver request handler
type HandlerFunc func(http.ResponseWriter, *http.Request) error

// GuardFn is guard middleware handler
type GuardFn func(r *http.Request) error

// WrapHandler wraps handler
func WrapHandler(h HandlerFunc, guards ...GuardFn) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func DeprecatedEndpoint(h HandlerFunc, sunsetDate time.Time) HandlerFunc {
	_ = "STUB: not implemented"
	return *new(HandlerFunc)
}

func handleError(err error, w http.ResponseWriter) { _ = "STUB: not implemented"; return }

// Ignore broken pipe errors
