package main

import (
	"os"

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
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	logger := log.New()

	cfg, err := config.Load()
	if err != nil {
		logger.Error().Err(err).Msg("load config failed")
		return err
	}

	if err := cfg.Validate(); err != nil {
		logger.Error().Err(err).Msg("invalid configuration")
		return err
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

	logger.Info().
		Any(logkeys.Env, cfg.Env).
		Str(logkeys.GitHash, CommitHash).
		Str(logkeys.BuildTime, BuildTime).
		Msg("application initializing...")

	application, cleanup, err := app.New(cfg, logger)
	if err != nil {
		logger.Error().Err(err).Msg("failed to init app")
		return err
	}
	defer cleanup()

	appCtx, cancel := ctx.GetAppCtx()
	defer cancel()

	logger.Info().Msg("starting application...")
	if err := application.Run(appCtx); err != nil {
		logger.Error().Err(err).Msg("application runtime error")
		return err
	}

	return nil
}
