package builder

import (
	"bytes"
	"context"
)

// BuildError is build error
type BuildError struct {
	message string
}

// Error implements error
func (e *BuildError) Error() string { _ = "STUB: not implemented"; return "" }

func newBuildError(msg string, args ...any) *BuildError { _ = "STUB: not implemented"; return nil }

func IsBuildError(err error) bool { _ = "STUB: not implemented"; return false }

func checkContextErrors(err error) (error, bool) { _ = "STUB: not implemented"; return nil, false }

func formatBuildError(ctx context.Context, err error, buff *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}
