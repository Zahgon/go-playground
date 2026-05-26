package goplay

import (
	"bytes"
	"io"
)

// FileSet is a helper to construct a Go playground request from multiple Go files.
type FileSet struct {
	goSourceFiles  map[string][]byte
	otherFiles     map[string][]byte
	goFileOrder    []string
	otherFileOrder []string
	buf            *bytes.Buffer
	dirty          bool
}

func NewFileSet(bufSize int) *FileSet { _ = "STUB: not implemented"; return nil }

// HasGoFiles returns whether a set contains at-least one ".go" file.
func (f *FileSet) HasGoFiles() bool { _ = "STUB: not implemented"; return false }

// Add adds a file to the buffer.
func (f *FileSet) Add(name string, src []byte) error { _ = "STUB: not implemented"; return nil }

func txtarAppendFile(buf *bytes.Buffer, fname string, data []byte) {
	_ = "STUB: not implemented"
	return
}

// If contents doesn't end with line break - add it.
// Required, as line break is txtar file separator.

func (f *FileSet) buildBuf() *bytes.Buffer {
	_ = "STUB: not implemented"

	// Skip if buffer is populated.
	return nil
}

// First, write Go source files and then other files.
// Upstream might misbehave if non-Go files come first.

func (f *FileSet) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (f *FileSet) Reader() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }
