package builder

import (
	"context"
	"time"

	"github.com/tevino/abool"
	"go.uber.org/zap"
)

type Cleaner interface {
	CleanJobName() string
	Clean(ctx context.Context) error
}

// CleanupDispatchService calls cleanup entries after periodical interval of time.
type CleanupDispatchService struct {
	isRunning abool.AtomicBool
	logger    *zap.Logger
	interval  time.Duration
	cleaners  []Cleaner
}

func NewCleanupDispatchService(logger *zap.Logger, interval time.Duration, cleaners ...Cleaner) *CleanupDispatchService {
	_ = "STUB: not implemented"
	return nil
}

func (c *CleanupDispatchService) Start(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *CleanupDispatchService) dispatchCleanup(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}
