package main

import (
	"time"

	"github.com/shanth1/gotools/consts"
	"github.com/shanth1/gotools/ctx"
	"github.com/shanth1/gotools/log"
	"github.com/shanth1/gotools/logkeys"
	"github.com/shanth1/template/internal/app"
	"github.com/shanth1/template/internal/config"
)

var (
	CommitHash = "n/a"
	BuildTime  = "n/a"
)

// @title           Golang Project Template API
// @version         1.0
// @description     This is a sample server following Hexagonal Architecture.
// @termsOfService  http://swagger.io/terms/

// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io

// @license.name    MIT
// @license.url     https://opensource.org/licenses/MIT

// @host            localhost:8080
// @BasePath        /
func main() {
	ctx, shutdownCtx, cancel, shutdownCancel := ctx.WithGracefulShutdown(10 * time.Second)
	defer cancel()
	defer shutdownCancel()

	logger := log.New()
	logger.Info().
		Str(logkeys.GitHash, CommitHash).
		Str(logkeys.BuildTime, BuildTime).
		Msg("starting service")

	cfg, err := config.Load()
	if err != nil {
		logger.Fatal().Err(err).Msg("load config")
	}

	if err := cfg.Validate(); err != nil {
		logger.Fatal().Err(err).Msg("invalid configuration")
	}

	logger = logger.WithOptions(log.WithConfig(log.Config{
		Level:        cfg.Logger.Level,
		App:          cfg.Logger.App,
		Service:      cfg.Logger.Service,
		UDPAddress:   cfg.Logger.UDPAddress,
		EnableCaller: cfg.Logger.EnableCaller,
		Console:      cfg.Env != consts.EnvProd,
		JSONOutput:   cfg.Env == consts.EnvProd,
	}))

	logger.Info().Any(logkeys.Env, cfg.Env).Msg("application has been successfully configured")

	ctx = log.NewContext(ctx, logger)
	app.Run(ctx, shutdownCtx, cfg)
}
