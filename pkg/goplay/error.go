package goplay

import (
	"net/http"
)

type HTTPError struct {
	StatusCode int
	Status     string
}

func NewHTTPError(rsp *http.Response) *HTTPError { _ = "STUB: not implemented"; return nil }

func (err HTTPError) Error() string { _ = "STUB: not implemented"; return "" }

// IsHTTPError checks if passed error is HTTPError.
func IsHTTPError(err error) (*HTTPError, bool) { _ = "STUB: not implemented"; return nil, false }
