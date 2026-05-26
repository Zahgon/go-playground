package builder

import (
	"regexp"
	"runtime"
)

var (
	goVersion      string
	goVersionRegEx = regexp.MustCompile(`(?m)^go (\d+\.\d+\.\d+){1}`)
)

func init() {
	goVersion = parseGoVersion(runtime.Version())
	if goVersion == "" {
		goVersion = "1.25"
	}
}

func parseGoVersion(input string) string { _ = "STUB: not implemented"; return "" }

func generateGoMod(modName string) []byte { _ = "STUB: not implemented"; return nil }
