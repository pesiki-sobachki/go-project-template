package app

import (
	"context"

	"github.com/shanth1/gotools/log"
	"github.com/shanth1/template/internal/config"
	"github.com/shanth1/template/internal/core/port"
	"github.com/shanth1/template/internal/core/service"
)

func Run(ctx, shutdownCtx context.Context, cfg *config.Config) {
	logger := log.FromContext(ctx)

	// TODO: outbound adapters

	service := service.New(logger)

	runHTTPServer(ctx, shutdownCtx, cfg, service, logger)

}

func runHTTPServer(
	ctx context.Context,
	_ context.Context, // TODO: shutdownCtx
	cfg *config.Config,
	_ port.Service, // TODO: service
	logger log.Logger,
) {
	// TODO: inbound adapters

	go func() {
		logger.Info().Msgf("starting HTTP server on %s", cfg.Addr)

		// TODO: ListenAndServe
		// if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// 	logger.Fatal().Err(err).Msg("http server failed")
		// }
	}()

	<-ctx.Done()
	logger.Info().Msg("shutting down HTTP server...")
	// TODO: shutdown
	// if err := server.Shutdown(shutdownCtx); err != nil {
	// 	logger.Error().Err(err).Msg("http server graceful shutdown failed")
	// }
}
