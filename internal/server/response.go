package server

import (
	"net/http"
)

const maxFilesCount = 10

type FilesPayload struct {
	Files map[string]string `json:"files"`
}

// Validate checks file name and contents and returns error on validation failure.
func (p FilesPayload) Validate() error { _ = "STUB: not implemented"; return nil }

// HasUnitTests checks whether file list contains any unit test.
//
// Note: at the moment, func doesn't check file contents and only check file names.
func (p FilesPayload) HasUnitTests() bool { _ = "STUB: not implemented"; return false }

// ErrorResponse is error response
type ErrorResponse struct {
	code int

	// Error is error message
	Error string `json:"error"`
}

// NewErrorResponse is ErrorResponse constructor
func NewErrorResponse(err error) *ErrorResponse { _ = "STUB: not implemented"; return nil }

// Write writes error to response
func (r *ErrorResponse) Write(w http.ResponseWriter) http.ResponseWriter {
	_ = "STUB: not implemented"
	return *new(http.ResponseWriter)
}

// WriteJSON encodes object as JSON and writes it to stdout
func WriteJSON(w http.ResponseWriter, i interface{}) { _ = "STUB: not implemented"; return }
