package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"golang.org/x/time/rate"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/x1unix/go-playground/internal/builder"
	"github.com/x1unix/go-playground/pkg/goplay"
)

var ErrEmptyRequest = errors.New("empty request")

type APIv2HandlerConfig struct {
	Client       *goplay.Client
	Builder      builder.BuildService
	BuildTimeout time.Duration
}

func (cfg APIv2HandlerConfig) buildContext(parentCtx context.Context) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

type APIv2Handler struct {
	logger  *zap.Logger
	limiter *rate.Limiter
	cfg     APIv2HandlerConfig
}

func NewAPIv2Handler(cfg APIv2HandlerConfig) *APIv2Handler { _ = "STUB: not implemented"; return nil }

// HandleGetSnippet handles requests to get snippet by id.
func (h *APIv2Handler) HandleGetSnippet(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Serve stuff as-is

// HandleShare handles snippet share requests.
func (h *APIv2Handler) HandleShare(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// HandleFormat handles gofmt requests.
func (h *APIv2Handler) HandleFormat(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// HandleRun sends snippet to upstream play.go.dev and returns evaluation result.
func (h *APIv2Handler) HandleRun(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// HandleCompile handles WebAssembly compile requests.
func (h *APIv2Handler) HandleCompile(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	// Limit for request timeout
	return nil
}

// Wait for our queue in line for compilation

func (h *APIv2Handler) Mount(r *mux.Router) { _ = "STUB: not implemented"; return }
