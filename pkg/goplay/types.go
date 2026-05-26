package goplay

import (
	"net/url"
	"time"
)

const DefaultVersion = 2

type Backend = string

const (
	BackendGoCurrent = ""
	BackendGoPrev    = "goprev"
	BackendGoTip     = "gotip"
)

// Snippet represents shared snippet
type Snippet struct {
	FileName string
	Contents string
}

// FmtResponse is the response returned from
// upstream play.golang.org/fmt request
type FmtResponse struct {
	Body  string
	Error string
}

// HasError returns error if any occurred
func (r *FmtResponse) HasError() error { _ = "STUB: not implemented"; return nil }

type CompileRequest struct {
	Version int
	WithVet bool
	Body    []byte
}

func (r CompileRequest) URLValues() url.Values { _ = "STUB: not implemented"; return *new(url.Values) }

// CompileEvent represents individual
// event record in CompileResponse
type CompileEvent struct {
	Message string
	Kind    string
	Delay   time.Duration
}

// CompileResponse is the response returned from
// upstream play.golang.org/compile request
type CompileResponse struct {
	Body   *string
	Events []*CompileEvent
	Errors string
}

// GetBody returns response body
func (cr CompileResponse) GetBody() string { _ = "STUB: not implemented"; return "" }

// HasError returns error if any occurred
func (cr *CompileResponse) HasError() error { _ = "STUB: not implemented"; return nil }
