package cmd

import (
	"encoding/json"
	"io"
)

func writeOutput(flags importsFlags, data any) error { _ = "STUB: not implemented"; return nil }

func getEncoder(dst io.Writer, pretty bool) *json.Encoder { _ = "STUB: not implemented"; return nil }

func silentClose(c io.Closer) {
	_ = "STUB: not implemented"
	// I don't care
	return
}
