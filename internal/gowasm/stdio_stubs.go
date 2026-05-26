//go:build !js

package gowasm

func wasmConsoleWrite(fd int, data []byte) { _ = "STUB: not implemented"; return }
