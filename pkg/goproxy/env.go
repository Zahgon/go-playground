package goproxy

func proxyAddressFromEnv() string { _ = "STUB: not implemented"; return "" }

// NewClientFromEnv returns a new client with address from GOPROXY environment variable.
//
// Attention: only first address is respected. "direct" value is also ignored.
//
// The default Go proxy URL is returned in case of error.
func NewClientFromEnv() *Client { _ = "STUB: not implemented"; return nil }
