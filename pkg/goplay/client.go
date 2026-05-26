// Package goplay provides a client for the official Go Playground API.
package goplay

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	DefaultUserAgent     = "goplay.tools/1.0 (http://goplay.tools/)"
	DefaultPlaygroundURL = "https://go.dev/_"

	// MaxSnippetSize value taken from
	// https://github.com/golang/playground/blob/master/app/goplay/share.go
	MaxSnippetSize = 64 * 1024
)

// Client is Go Playground API client
type Client struct {
	client    http.Client
	baseUrl   string
	userAgent string
}

// NewClient returns new Go playground client
func NewClient(baseUrl, userAgent string, timeout time.Duration) *Client {
	_ = "STUB: not implemented"
	return nil
}

// NewDefaultClient returns Go Playground client with defaults
func NewDefaultClient() *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) newRequest(ctx context.Context, method, queryPath string, body io.Reader) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) getRequest(ctx context.Context, queryPath string) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) doRequest(ctx context.Context, method, url, contentType string, body io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) postForm(ctx context.Context, url string, data url.Values) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) postJSON(ctx context.Context, url string, data url.Values, out interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
