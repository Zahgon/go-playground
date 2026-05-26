package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/x1unix/go-playground/internal/announcements"
	"github.com/x1unix/go-playground/internal/builder"
	"github.com/x1unix/go-playground/internal/server/backendinfo"
	"github.com/x1unix/go-playground/pkg/goplay"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

const (
	// limit for wasm compile requests per second (250ms per request)
	compileRequestsPerFrame = 4
	frameTime               = time.Second

	wasmMimeType     = "application/wasm"
	artifactParamVal = "artifactId"
)

// APIv1Handler is API v1 handler
type APIv1Handler struct {
	config          ServiceConfig
	log             *zap.SugaredLogger
	compiler        builder.BuildService
	versionProvider backendinfo.BackendVersionProvider

	client  *goplay.Client
	limiter *rate.Limiter
}

type ServiceConfig struct {
	Version      string
	Announcement *announcements.Announcement
}

// NewAPIv1Handler is APIv1Handler constructor
func NewAPIv1Handler(cfg ServiceConfig, client *goplay.Client, builder builder.BuildService, versionProvider backendinfo.BackendVersionProvider) *APIv1Handler {
	_ = "STUB: not implemented"
	return nil
}

// Mount mounts service on route
func (s *APIv1Handler) Mount(r *mux.Router) { _ = "STUB: not implemented"; return }

// HandleGetVersion handles /api/version
func (s *APIv1Handler) HandleGetVersion(w http.ResponseWriter, _ *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// HandleGetAnnouncement returns service announcement banner contents.
func (s *APIv1Handler) HandleGetAnnouncement(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *APIv1Handler) HandleGetVersions(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// HandleArtifactRequest handles WASM build artifact request
func (s *APIv1Handler) HandleArtifactRequest(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}
