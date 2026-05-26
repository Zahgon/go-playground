package server

import (
	"net/url"

	"github.com/x1unix/go-playground/pkg/goplay"
)

// RunParams is code run request parameters
type RunParams struct {
	// Vet enables go vet
	Vet bool

	// Format identifies whether to format code before run.
	//
	// Deprecated and used only for v1.
	Format bool

	// Backend is Go run backend.
	Backend string
}

func RunParamsFromQuery(query url.Values) (params RunParams, err error) {
	_ = "STUB: not implemented"
	return *new(RunParams), nil
}

func parseBoolQueryParam(query url.Values, key string, defaults bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func isContentLengthError(err error) bool { _ = "STUB: not implemented"; return false }

func backendFromQuery(query url.Values) (goplay.Backend, error) {
	_ = "STUB: not implemented"
	return *new(goplay.Backend), nil
}
