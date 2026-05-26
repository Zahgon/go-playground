// Package goproxy provides Go modules proxy API client.
package goproxy

import (
	"context"
	"io"
	"net/http"
)

const (
	jsonContentType = "application/json"

	DefaultProxyURL = "https://proxy.golang.org"
)

type Client struct {
	url    string
	client *http.Client
}

func NewClient(client *http.Client, baseURL string) *Client { _ = "STUB: not implemented"; return nil }

// NewClientWithDefaults returns a new go modules proxy client with default http client
// and official Go Modules proxy server.
func NewClientWithDefaults() *Client { _ = "STUB: not implemented"; return nil }

// SetBaseURL updates proxy base URL
func (c *Client) SetBaseURL(baseURL string) { _ = "STUB: not implemented"; return }

func (c *Client) getURL(segments ...string) string { _ = "STUB: not implemented"; return "" }

func (c *Client) get(ctx context.Context, segments ...string) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) getBody(ctx context.Context, segments ...string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (c *Client) getJSON(ctx context.Context, out any, segments ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetVersions returns a list of module versions.
func (c *Client) GetVersions(ctx context.Context, pkgUrl string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use bytes to avoid multiple string copy

// GetLatestVersion returns information about a latest module version.
func (c *Client) GetLatestVersion(ctx context.Context, pkgUrl string) (*VersionInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetVersionInfo returns information about module version.
func (c *Client) GetVersionInfo(ctx context.Context, pkgUrl, version string) (*VersionInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetModuleFile returns go.mod file of specific version of a module.
func (c *Client) GetModuleFile(ctx context.Context, pkgUrl, version string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetModuleSource returns zip archive stream for that version of the given module.
//
// Returns a io.ReadCloser-compatible struct with size and content type information.
func (c *Client) GetModuleSource(ctx context.Context, pkgUrl, version string) (*ArchiveReadCloser, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func silentClose(closer io.Closer) { _ = "STUB: not implemented"; return }
