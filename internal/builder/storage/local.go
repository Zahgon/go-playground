package storage

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/tevino/abool"
	"go.uber.org/zap"
)

const (
	srcDirName  = "src"
	binDirName  = "bin"
	workDirName = "goplay-builds"

	maxCleanTime = time.Second * 10
	perm         = 0744
)

type cachedFile struct {
	io.ReadCloser

	size    int64
	useLock *sync.Mutex
}

func (c cachedFile) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (c cachedFile) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// LocalStorage is local build artifact storage
type LocalStorage struct {
	log     *zap.Logger
	useLock *sync.Mutex
	dirty   *abool.AtomicBool
	gcRun   *abool.AtomicBool
	workDir string
	srcDir  string
	binDir  string
}

// NewLocalStorage constructs new local storage
func NewLocalStorage(log *zap.Logger, baseDir string) (ls *LocalStorage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isDirDirty(dir string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (s LocalStorage) getOutputLocation(id ArtifactID) string { _ = "STUB: not implemented"; return "" }

// HasItem implements storage interface
func (s LocalStorage) HasItem(id ArtifactID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetItem implements storage interface
func (s LocalStorage) GetItem(id ArtifactID) (ReadCloseSizer, error) {
	_ = "STUB: not implemented"
	return *new(ReadCloseSizer), nil
}

// CreateWorkspace implements storage interface
func (s LocalStorage) CreateWorkspace(id ArtifactID, files map[string][]byte) (*Workspace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// mark storage as dirty

// Ensure bin dir exists

// Write entries

func (s LocalStorage) clean() error { _ = "STUB: not implemented"; return nil }

// cleanup sources and binaries

// remove dirty flag

// CleanJobName implements builder.Cleaner interface.
func (s LocalStorage) CleanJobName() string {
	_ = "STUB: not implemented"

	// Clean implements' builder.Cleaner interface.
	return ""
}

func (s LocalStorage) Clean(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func createParentDir(workDir, fileName string) error { _ = "STUB: not implemented"; return nil }
