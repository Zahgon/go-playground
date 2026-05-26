package server

import (
	"errors"
	"net/http"

	"github.com/x1unix/go-playground/pkg/goplay"
)

// evalPayloadFromRequest validates and extracts snippet payload for Go Playground API evaluate request.
func evalPayloadFromRequest(r *http.Request) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Go playground might behave inadequately with unit-tests when multi-file snippet has only one file.
// See: https://github.com/x1unix/go-playground/issues/324

// Official Go Playground can't run unit-tests in multi-file snippets.
// We can't do anything about it, just detect a case and return an error to a user.
// See: https://github.com/golang/go/issues/68327

var errNoGoFiles = NewBadRequestError(
	errors.New("no Go files"),
)

func fileSetFromRequest(r *http.Request) (*goplay.FileSet, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func buildFilesFromRequest(r *http.Request) (map[string][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filesPayloadFromRequest(r *http.Request) (*FilesPayload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
