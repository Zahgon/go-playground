package server

import (
	"io"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"
)

type TemplateArguments struct {
	GoogleTagID string
}

type TemplateFileServer struct {
	log          *zap.Logger
	filePath     string
	templateVars TemplateArguments
	once         *sync.Once
	buffer       io.ReadSeeker
	modTime      time.Time
}

// NewTemplateFileServer returns handler which compiles and serves HTML page template.
func NewTemplateFileServer(logger *zap.Logger, filePath string, tplVars TemplateArguments) *TemplateFileServer {
	_ = "STUB: not implemented"
	return nil
}

func (fs *TemplateFileServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (fs *TemplateFileServer) precompileTemplate() { _ = "STUB: not implemented"; return }
