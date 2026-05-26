package goplay

import "errors"

// ErrSnippetNotFound is snippet not found error
var ErrSnippetNotFound = errors.New("snippet not found")

// CompileFailedError is build error
type CompileFailedError struct {
	msg string
}

// Error implements error
func (c CompileFailedError) Error() string {
	_ = "STUB: not implemented"

	// IsCompileError checks if error is CompileFailedError
	return ""
}

func IsCompileError(err error) bool { _ = "STUB: not implemented"; return false }
