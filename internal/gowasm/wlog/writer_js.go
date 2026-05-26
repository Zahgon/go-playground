package wlog

import (
	"io"
)

var (
	// StdLog is standard output for worker logs.
	StdLog io.Writer = newLogWriter(logLevelInfo)

	// StdDebug is debug output for worker logs.
	StdDebug io.Writer = newLogWriter(logLevelDebug)
)

//go:wasmimport gojs github.com/x1unix/go-playground/internal/gowasm/wlog.logWrite
func logWrite(level uint8, data []byte)

const (
	logLevelDebug = 0
	logLevelInfo  = 1
)

type logWriter struct {
	level uint8
}

func newLogWriter(level uint8) logWriter { _ = "STUB: not implemented"; return *new(logWriter) }

func (w logWriter) Write(data []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
