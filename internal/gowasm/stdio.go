package gowasm

import (
	"syscall"
)

var (
	Stdout = newConsoleWriter(syscall.Stdout)
	Stderr = newConsoleWriter(syscall.Stderr)
)

type consoleWriter struct {
	fdId int
}

func newConsoleWriter(fdId int) consoleWriter {
	_ = "STUB: not implemented"
	return *new(consoleWriter)
}

func (w consoleWriter) Write(data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
