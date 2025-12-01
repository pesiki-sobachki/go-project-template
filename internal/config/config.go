package config

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	Env    string     `mapstructure:"env" yaml:"env" validate:"required,oneof=local development production"`
	Addr   string     `mapstructure:"addr" yaml:"addr" validate:"required,hostname_port"`
	HTTP   HTTPConfig `mapstructure:"http" yaml:"http" validate:"required"`
	Logger Logger     `mapstructure:"logger" yaml:"logger" validate:"required"`
}

type HTTPConfig struct {
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

func (c *Config) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
