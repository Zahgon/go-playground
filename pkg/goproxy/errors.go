package goproxy

import (
	"net/http"
)

type HTTPError struct {
	Code    int
	URL     string
	Message string
}

func buildHTTPError(req *http.Request, rsp *http.Response) *HTTPError {
	_ = "STUB: not implemented"
	return nil
}

func (err *HTTPError) Error() string { _ = "STUB: not implemented"; return "" }

func IsHTTPError(err error) (*HTTPError, bool) { _ = "STUB: not implemented"; return nil, false }

func IsNotFoundError(err error) (*HTTPError, bool) { _ = "STUB: not implemented"; return nil, false }
