package builder

import (
	"context"
	"os"

	"go.uber.org/zap"

	"github.com/x1unix/go-playground/internal/builder/storage"
	"github.com/x1unix/go-playground/pkg/util/osutil"
)

// defaultGoModName is default module name that will be set if no go.mod provided.
const defaultGoModName = "app"

// predefinedBuildVars is list of environment vars which contain build values
var predefinedBuildVars = osutil.EnvironmentVariables{
	"CGO_ENABLED": "0",
	"GOOS":        "js",
	"GOARCH":      "wasm",
	"HOME":        os.Getenv("HOME"),
}

// Result is WASM build result
type Result struct {
	// FileName is artifact file name
	FileName string

	// IsTest indicates whether binary is a test file
	IsTest bool

	// HasBenchmark indicates whether test contains benchmarks.
	HasBenchmark bool

	// HasFuzz indicates whether test has fuzzing tests.
	HasFuzz bool
}

// BuildEnvironmentConfig is BuildService environment configuration.
type BuildEnvironmentConfig struct {
	// IncludedEnvironmentVariables is a list included environment variables for build.
	IncludedEnvironmentVariables osutil.EnvironmentVariables

	// KeepGoModCache disables Go modules cache cleanup.
	KeepGoModCache bool
}

// BuildService is WASM build service
type BuildService struct {
	log       *zap.Logger
	config    BuildEnvironmentConfig
	storage   storage.StoreProvider
	cmdRunner CommandRunner
}

// NewBuildService is BuildService constructor
func NewBuildService(log *zap.Logger, cfg BuildEnvironmentConfig, store storage.StoreProvider) BuildService {
	_ = "STUB: not implemented"
	return *new(BuildService)
}

func (s BuildService) getEnvironmentVariables() []string { _ = "STUB: not implemented"; return nil }

// GetArtifact returns artifact by id
func (s BuildService) GetArtifact(id storage.ArtifactID) (storage.ReadCloseSizer, error) {
	_ = "STUB: not implemented"
	return *new(storage.ReadCloseSizer), nil
}

// Build compiles Go source to WASM and returns result
func (s BuildService) Build(ctx context.Context, files map[string][]byte) (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Go module is required to build project

// Just return precompiled result if data is cached already

// Immediately schedule cleanup job!

func (s BuildService) buildSource(ctx context.Context, projInfo projectInfo, workspace *storage.Workspace) error {
	_ = "STUB: not implemented"
	// Populate go.mod and go.sum files.
	return nil
}

func (s BuildService) handleNoSpaceLeft() { _ = "STUB: not implemented"; return }

func (s BuildService) runGoTool(ctx context.Context, workDir string, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// CleanJobName implements' builder.Cleaner interface.
func (s BuildService) CleanJobName() string {
	_ = "STUB: not implemented"

	// Clean implements' builder.Cleaner interface.
	//
	// Cleans go build and modules cache.
	return ""
}

func (s BuildService) Clean(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
