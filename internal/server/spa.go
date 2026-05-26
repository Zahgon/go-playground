package server

import (
	"net/http"
)

const (
	IndexFileName    = "index.html"
	NotFoundFileName = "404.html"
)

type httpStatusInterceptor struct {
	http.ResponseWriter
	desiredStatus int
}

func (i httpStatusInterceptor) WriteHeader(_ int) { _ = "STUB: not implemented"; return }

// SpaFileServer is a wrapper around http.FileServer for serving SPA contents.
type SpaFileServer struct {
	root            string
	NotFoundHandler http.Handler
	templateVars    TemplateArguments
}

// ServeHTTP implements http.Handler
func (fs *SpaFileServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

//if empty, set current directory

//add prefix and clean

//path to file

//check if file exists

func containsDotDot(v string) bool { _ = "STUB: not implemented"; return false }

func isSlashRune(r rune) bool { _ = "STUB: not implemented"; return false }

// NewSpaFileServer returns SPA handler
func NewSpaFileServer(root string, tplVars TemplateArguments) *SpaFileServer {
	_ = "STUB: not implemented"
	return nil
}

// NewFileServerWithStatus returns http.Handler which serves specified file with desired HTTP status
func NewFileServerWithStatus(name string, code int) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// ServeFileWithStatus serves file in HTTP response with specified HTTP status.
func ServeFileWithStatus(rw http.ResponseWriter, r *http.Request, name string, code int) {
	_ = "STUB: not implemented"
	return
}
