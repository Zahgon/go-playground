package main

import (
	"context"
	"net/http"
	"sync"

	"github.com/x1unix/go-playground/internal/config"
	"github.com/x1unix/go-playground/pkg/util/cmdutil"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

// Version is server version symbol. Should be replaced by linker during build
var Version = "testing"

func main() {
	cfg, err := config.FromEnv(config.FromFlags())
	if err != nil {
		cmdutil.FatalOnError(err)
	}

	logger, err := cfg.Log.ZapLogger()
	if err != nil {
		cmdutil.FatalOnError(err)
	}
	zap.ReplaceGlobals(logger)
	defer logger.Sync() //nolint:errcheck

	if err := cfg.Validate(); err != nil {
		logger.Fatal("invalid server configuration", zap.Error(err))
	}

	if err := start(logger, cfg); err != nil {
		logger.Fatal("Failed to start application", zap.Error(err))
	}
}

func start(logger *zap.Logger, cfg *config.Config) error { _ = "STUB: not implemented"; return nil }

// Initialize services

// Start cleanup service

// Initialize API endpoints

// Web UI routes

func startHttpServer(ctx context.Context, wg *sync.WaitGroup, server *http.Server) error {
	_ = "STUB: not implemented"
	return nil
}
