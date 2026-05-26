package server

import (
	"net/http"

	"github.com/x1unix/go-playground/pkg/goplay"
)

// ErrSnippetTooLarge is snippet max size limit error
var ErrSnippetTooLarge = Errorf(
	http.StatusRequestEntityTooLarge,
	"code snippet too large (max %d bytes)",
	goplay.MaxSnippetSize,
)

// HTTPError is HTTP response error
type HTTPError struct {
	code   int
	parent error
}

// Error implements error
func (err *HTTPError) Error() string { _ = "STUB: not implemented"; return "" }

// Unwrap implements error
func (err *HTTPError) Unwrap() error {
	_ = "STUB: not implemented"

	// WriteResponse writes error to response
	return nil
}

func (err *HTTPError) WriteResponse(rw http.ResponseWriter) { _ = "STUB: not implemented"; return }

// NewHTTPError constructs a new error
func NewHTTPError(code int, err error) *HTTPError { _ = "STUB: not implemented"; return nil }

// Errorf returns new formatted error
func Errorf(code int, format string, args ...interface{}) *HTTPError {
	_ = "STUB: not implemented"
	return nil
}

func NewBadRequestError(err error) *HTTPError { _ = "STUB: not implemented"; return nil }
