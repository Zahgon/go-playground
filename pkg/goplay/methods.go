package goplay

import (
	"context"
	"io"
)

// GetSnippet returns snippet from Go playground
func (c *Client) GetSnippet(ctx context.Context, snippetID string) (*Snippet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Share shares snippet to go playground
func (c *Client) Share(ctx context.Context, src io.Reader) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GoImports performs Goimports
func (c *Client) GoImports(ctx context.Context, src []byte, backend Backend) (*FmtResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Evaluate runs code in goplayground and returns response
func (c *Client) Evaluate(ctx context.Context, req CompileRequest, backend string) (*CompileResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateBackend validates Go Playground backend name.
func ValidateBackend(backend Backend) bool { _ = "STUB: not implemented"; return false }

func appendBackendToPath(urlPath string, backend Backend) string {
	_ = "STUB: not implemented"
	return ""
}
