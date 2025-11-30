package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/shanth1/gotools/conf"
	"github.com/shanth1/gotools/consts"
	"github.com/shanth1/gotools/ctx"
	"github.com/shanth1/gotools/flags"
	"github.com/shanth1/gotools/log"
	"github.com/shanth1/template/internal/app"
	"github.com/shanth1/template/internal/config"
)

var (
	CommitHash = "n/a"
	BuildTime  = "n/a"
)

type Flags struct {
	ConfigPath string `flag:"config" usage:"Path to the YAML config file"`
}

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
	fmt.Printf("Starting Service\nCommit: %s\nBuild Time: %s\n", CommitHash, BuildTime)

	ctx, shutdownCtx, cancel, shutdownCancel := ctx.WithGracefulShutdown(10 * time.Second)
	defer cancel()
	defer shutdownCancel()

	logger := log.New()

	flagCfg := &Flags{}
	if err := flags.RegisterFromStruct(flagCfg); err != nil {
		logger.Fatal().Err(err).Msg("register flags")
	}
	flag.Parse()

	cfg := &config.Config{}
	if err := conf.Load(flagCfg.ConfigPath, cfg); err != nil {
		logger.Fatal().Err(err).Msg("load config")
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

	ctx = log.NewContext(ctx, logger)
	app.Run(ctx, shutdownCtx, cfg)
}
