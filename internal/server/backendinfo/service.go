package backendinfo

import (
	"context"
	_ "embed"
	"time"

	"github.com/x1unix/go-playground/pkg/goplay"
	"go.uber.org/zap"
)

const (
	goVersionRetryAttempts = 3
	goVersionRetryDelay    = time.Second

	DefaultVersionCacheTTL = 48 * time.Hour
)

//go:embed resources/version.go.txt
var versionSnippet []byte

const cacheFileVersion = 1

var _ BackendVersionProvider = (*BackendVersionService)(nil)

type ServiceConfig struct {
	// Version is cache file version
	Version int

	// CacheFile is name of a file which will be used to cache Go playground versions.
	CacheFile string

	// TTL is expiration interval.
	TTL time.Duration
}

type cacheEntry struct {
	Version   int
	CreatedAt time.Time
	Data      BackendVersions
}

// BackendVersionService provides information about used Go versions
// for all backends.
type BackendVersionService struct {
	logger *zap.Logger
	client *goplay.Client
	cfg    ServiceConfig

	memCache *cacheEntry
}

func NewBackendVersionService(logger *zap.Logger, client *goplay.Client, cfg ServiceConfig) *BackendVersionService {
	_ = "STUB: not implemented"
	return nil
}

func (svc *BackendVersionService) ServerVersion() string { _ = "STUB: not implemented"; return "" }

func (svc *BackendVersionService) visitCache() (*cacheEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetVersions provides Go version information for all backends.
func (svc *BackendVersionService) GetRemoteVersions(ctx context.Context) (*BackendVersions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (svc *BackendVersionService) populateVersionCache(ctx context.Context) (*BackendVersions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (svc *BackendVersionService) cacheVersions(versions *BackendVersions) error {
	_ = "STUB: not implemented"
	return nil
}

func (svc *BackendVersionService) pullBackendVersions(ctx context.Context) (*BackendVersions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Playground "gotip" and "goprev" backends are often broken
// and I'm getting tired of seeing 5xx responses if just one of them is dead.
//
// Throw only if stable version is down. For others - try to figure out fallback values.

// We don't afraid race condition because each backend is written to a separate address

func (svc *BackendVersionService) fetchGoBackendVersionWithRetry(ctx context.Context, backend goplay.Backend) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Retry only on server issues

func (svc *BackendVersionService) getGoBackendVersion(ctx context.Context, backend goplay.Backend) (string, error) {
	_ = "STUB: not implemented"
	// Dirty hack to fetch Go version for playground backend by running a simple program
	// which returns Go version to stdout.
	return "", nil
}

func normalizeGoVersion(str string) string { _ = "STUB: not implemented"; return "" }
