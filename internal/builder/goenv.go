package builder

import (
	"context"
	"errors"
	"os/exec"
)

const goRootEnv = "GOROOT"

// ErrUndefinedEnvVariable occurs when requested environment variable is undefined or empty
var ErrUndefinedEnvVariable = errors.New("environment variable is undefined or empty")

// GOROOT returns host GOROOT variable from OS environment vars or from Go tool environment.
func GOROOT() (string, error) { _ = "STUB: not implemented"; return "", nil }

func newGoToolCommand(ctx context.Context, args ...string) *exec.Cmd {
	_ = "STUB: not implemented"
	return nil
}

func getEnvFromGo(ctx context.Context, envName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// LookupEnv gets variable by name from shell environment or using Go environment using "go env" tool.
func LookupEnv(ctx context.Context, varName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Lookup env variable using "go env" tool if not defined in environment.
