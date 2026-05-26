package config

import (
	"flag"
	"time"

	"github.com/x1unix/go-playground/internal/announcements"
)

const (
	DefaultWriteTimeout   = 60 * time.Second
	DefaultReadTimeout    = 15 * time.Second
	DefaultIdleTimeout    = 90 * time.Second
	DefaultGoBuildTimeout = 40 * time.Second
	DefaultCleanInterval  = 10 * time.Minute
)

type HTTPConfig struct {
	// Addr is HTTP server listen address
	Addr string `envconfig:"APP_HTTP_ADDR" json:"addr"`

	// AssetsDir is directory which contains frontend assets
	AssetsDir string `envconfig:"APP_ASSETS_DIR" json:"assetsDir"`

	// WriteTimeout is HTTP response write timeout.
	WriteTimeout time.Duration `envconfig:"HTTP_WRITE_TIMEOUT"`

	// ReadTimeout is HTTP request read timeout.
	ReadTimeout time.Duration `envconfig:"HTTP_READ_TIMEOUT"`

	// IdleTimeout is delay timeout between requests to keep connection alive.
	IdleTimeout time.Duration `envconfig:"HTTP_IDLE_TIMEOUT"`
}

func (cfg *HTTPConfig) mountFlagSet(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

type PlaygroundConfig struct {
	// PlaygroundURL is Go playground server URL
	PlaygroundURL string `envconfig:"APP_PLAYGROUND_URL" json:"playgroundUrl"`

	// ConnectTimeout is HTTP request timeout for playground requests
	ConnectTimeout time.Duration `envconfig:"APP_PLAYGROUND_TIMEOUT" json:"connectTimeout"`
}

func (cfg *PlaygroundConfig) mountFlagSet(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

type BuildConfig struct {
	// BuildDir is path to directory to cache WebAssembly builds
	BuildDir string `envconfig:"APP_BUILD_DIR" json:"buildDir"`

	// CleanupInterval is WebAssembly build artifact cache clean interval
	CleanupInterval time.Duration `envconfig:"APP_CLEAN_INTERVAL" json:"cleanupInterval"`

	// GoBuildTimeout is Go program build timeout.
	GoBuildTimeout time.Duration `envconfig:"APP_GO_BUILD_TIMEOUT" json:"goBuildTimeout"`

	// SkipModuleCleanup disables Go module cache cleanup.
	SkipModuleCleanup bool `envconfig:"APP_SKIP_MOD_CLEANUP" json:"skipModuleCleanup"`

	// BypassEnvVarsList is allow-list of environment variables
	// that will be passed to Go compiler.
	//
	// Empty value disables environment variable filter.
	BypassEnvVarsList []string `envconfig:"APP_PERMIT_ENV_VARS" json:"bypassEnvVarsList"`
}

func (cfg *BuildConfig) mountFlagSet(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

type ServicesConfig struct {
	// GoogleAnalyticsID is Google Analytics tag ID (optional)
	GoogleAnalyticsID string `envconfig:"APP_GTAG_ID" json:"googleAnalyticsID"`
}

func (cfg *ServicesConfig) mountFlagSet(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

type MiscConfig struct {
	// Announcement allows to set announcement banner displayed on a header banner.
	//
	// This can be used to display a server maintenance warning or other important info.
	//
	// Announcement string should be encoded using '/cmd/announcements' tool.
	Announcement announcements.TextMarshaler `envconfig:"SERVER_ANNOUNCEMENT"`
}

type Config struct {
	HTTP       HTTPConfig       `json:"http"`
	Playground PlaygroundConfig `json:"playground"`
	Build      BuildConfig      `json:"build"`
	Log        LogConfig        `json:"log"`
	Services   ServicesConfig   `json:"services"`
	Misc       MiscConfig       `json:"misc"`
}

// Validate validates a config and returns error if config is invalid.
func (cfg Config) Validate() error { _ = "STUB: not implemented"; return nil }

// FromFlagSet returns config file which will read values from flags
// when flag.Parse will be called.
func FromFlagSet(f *flag.FlagSet) *Config { _ = "STUB: not implemented"; return nil }

// FromFlags return config from parsed process arguments.
func FromFlags() *Config { _ = "STUB: not implemented"; return nil }

// FromEnv populates config with values from environment variables.
//
// If passed config is nil - a new config will be returned.
func FromEnv(input *Config) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }
