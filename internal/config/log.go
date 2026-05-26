package config

import (
	"flag"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type SentryConfig struct {
	DSN             string        `envconfig:"SENTRY_DSN" json:"dsn"`
	UseBreadcrumbs  bool          `envconfig:"SENTRY_USE_BREADCRUMBS" json:"useBreadcrumbs"`
	BreadcrumbLevel zapcore.Level `envconfig:"SENTRY_BREADCRUMB_LEVEL" json:"breadcrumbLevel"`
}

type LogConfig struct {
	Debug  bool          `envconfig:"APP_DEBUG" json:"debug"`
	Level  zapcore.Level `envconfig:"APP_LOG_LEVEL" json:"level"`
	Format string        `envconfig:"APP_LOG_FORMAT" json:"format"`

	Sentry SentryConfig `json:"sentry"`
}

func (cfg *LogConfig) mountFlagSet(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

// ZapLogger constructs a new zap.Logger instance from configuration.
func (cfg LogConfig) ZapLogger() (*zap.Logger, error) { _ = "STUB: not implemented"; return nil, nil }

// Return plain logger if sentry is disabled.

//when to send message to sentry
// enable sending breadcrumbs to Sentry
// at what level should we sent breadcrumbs to sentry

func (cfg LogConfig) getSentryDsn() string { _ = "STUB: not implemented"; return "" }
