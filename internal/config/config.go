package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/shanth1/gotools/conf"
	"github.com/shanth1/gotools/env"
	"github.com/shanth1/gotools/flags"
)

type Config struct {
	Env     string  `yaml:"-" env:"APP_ENV" validate:"required,oneof=local dev stage prod"`
	Addr    string  `mapstructure:"addr" yaml:"addr" env:"ADDR" validate:"required,hostname_port"`
	HTTP    HTTP    `mapstructure:"http" yaml:"http" validate:"required"`
	Logger  Logger  `mapstructure:"logger" yaml:"logger" validate:"required"`
	Metrics Metrics `mapstructure:"metrics" yaml:"metrics"`
}

type HTTP struct {
	ReadTimeout    time.Duration `mapstructure:"read_timeout" yaml:"read_timeout" validate:"min=100ms"`
	WriteTimeout   time.Duration `mapstructure:"write_timeout" yaml:"write_timeout" validate:"min=100ms"`
	IdleTimeout    time.Duration `mapstructure:"idle_timeout" yaml:"idle_timeout" validate:"min=1s"`
	RequestTimeout time.Duration `mapstructure:"request_timeout" yaml:"request_timeout" validate:"min=100ms"`
}

type Logger struct {
	App          string `mapstructure:"app" yaml:"app" validate:"required"`
	Level        string `mapstructure:"level" yaml:"level" validate:"required,oneof=debug info warn error fatal panic trace"`
	Service      string `mapstructure:"service" yaml:"service" validate:"required"`
	UDPAddress   string `mapstructure:"udp_address" yaml:"udp_address" validate:"omitempty,hostname_port"`
	EnableCaller bool   `mapstructure:"enable_caller" yaml:"enable_caller"`
}

type Metrics struct {
	User     string `env:"METRICS_USER"`
	Password string `env:"METRICS_PASSWORD"`
}

func (c *Config) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

type bootstrapConfig struct {
	AppEnv     string `flag:"env" usage:"Environment: local, dev, stage prod"`
	ConfigPath string `flag:"config" usage:"Path to the YAML config file"`
	EnvPath    string `flag:"env-path" usage:"Path to the env file"`
}

func Load() (*Config, error) {
	boot := &bootstrapConfig{}
	if err := flags.RegisterFromStruct(boot); err != nil {
		return nil, fmt.Errorf("register flags: %w", err)
	}
	flag.Parse()

	appEnv := boot.AppEnv
	if appEnv == "" {
		env, exists := os.LookupEnv("APP_ENV")
		if !exists {
			return nil, errors.New("app env param is empty")
		}
		appEnv = env
	}

	if boot.ConfigPath == "" {
		boot.ConfigPath = filepath.Join("config", fmt.Sprintf("config.%s.yaml", appEnv))
	}

	cfg := &Config{}
	if err := conf.Load(boot.ConfigPath, cfg); err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	if err := env.LoadIntoStruct(boot.EnvPath, cfg); err != nil {
		return nil, fmt.Errorf("load env: %w", err)
	}

	cfg.Env = appEnv
	return cfg, nil
}
