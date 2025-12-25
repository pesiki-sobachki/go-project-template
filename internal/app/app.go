package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/shanth1/gotools/log"
	transport "github.com/shanth1/template/internal/adapter/driving/http"
	"github.com/shanth1/template/internal/config"
	"github.com/shanth1/template/internal/core/service"
)

type App struct {
	cfg        *config.Config
	logger     log.Logger
	httpServer *http.Server
}

func New(cfg *config.Config, logger log.Logger) (*App, func(), error) {
	var cleanups []func()
	cleanup := func() {
		for i := len(cleanups) - 1; i >= 0; i-- {
			cleanups[i]()
		}
	}

	// TODO: driven adapters (repo, cache, filestorage, etc.)
	// TODO: driven adapters (close logic) -> cleanups

	// driven adapter -> service constructor
	svc := service.New(logger)
	httpHandler := transport.NewRouter(cfg, svc, logger)

	srv := &http.Server{
		Addr:              cfg.Addr,
		Handler:           httpHandler,
		ReadTimeout:       cfg.HTTP.ReadTimeout,
		WriteTimeout:      cfg.HTTP.WriteTimeout,
		IdleTimeout:       cfg.HTTP.IdleTimeout,
		ReadHeaderTimeout: 2 * time.Second,
	}

	return &App{
		cfg:        cfg,
		logger:     logger,
		httpServer: srv,
	}, cleanup, nil
}

func (a *App) Run(ctx context.Context) error {
	g, gCtx := errgroup.WithContext(ctx)

	// background workers are added in a similar way
	g.Go(func() error {
		a.logger.Info().Msgf("starting HTTP server on %s", a.cfg.Addr)
		if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("http server failed: %w", err)
		}
		return nil
	})

	<-gCtx.Done()

	a.logger.Info().Msg("shutting down application...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.httpServer.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("http server shutdown failed: %w", err)
	}

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	a.logger.Info().Msg("application stopped gracefully")
	return nil
}
