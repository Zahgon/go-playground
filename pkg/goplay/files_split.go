package goplay

import (
	"regexp"
)

var delimiterRegEx = regexp.MustCompile(`(?i)^-- (.*) --$`)

// supportedFileExtensions is a list of supported extensions except ".go" files.
var supportedFileExtensions = []string{
	".txt", ".json",
}

func cleanPath(pathname string) string {
	_ = "STUB: not implemented"
	// path.Clean don't clean dots for non-abs paths
	return ""
}

// ValidateFilePath validates a given file path contains a supported file.
//
// Filters out files that are not go.mod, *.go, *.txt or *.json files.
func ValidateFilePath(name string, strict bool) (isGoFile bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func isSeparatorLine(line string) (string, bool) { _ = "STUB: not implemented"; return "", false }

type SplitFileOpts struct {
	// DefaultFileName is a file name to use of source string doesn't contain any file name.
	DefaultFileName string

	// CheckPaths enables file path format validation.
	CheckPaths bool
}

// SplitFileSet splits string that contains source for multiple files in Go playground format.
//
// The official Go Playground defines a format to pass multiple files in request.
// Every file has to be separated with a special line which contains file name.
//
// For example:
//
//	package main
//
//	func main() {
//		...
//	}
//	-- foo/foo.go --
//	package foo
//
//	func Foo() {
//		...
//	}
func SplitFileSet(src string, opts SplitFileOpts) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip commit if string starts with file delimiter

// Commit previous chunk

// HACK: preserve trailing newline
