package testutil

import (
	"io"
	"testing"
)

// OpenTestdata returns opened file from testdata
func OpenTestdata(t *testing.T, fPath string) io.ReadWriteCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadWriteCloser)
}

// ReadTestdata reads testdata file
func ReadTestdata(t *testing.T, fPath string) []byte { _ = "STUB: not implemented"; return nil }

// TestdataPath returns file path from testdata
func TestdataPath(fPath string) string { _ = "STUB: not implemented"; return "" }
