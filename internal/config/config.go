package config

import "time"

type Config struct {
	Env    string     `mapstructure:"env" yaml:"env"`
	Addr   string     `mapstructure:"addr" yaml:"addr"`
	HTTP   HTTPConfig `mapstructure:"http" yaml:"http"`
	Logger Logger     `mapstructure:"logger" yaml:"logger"`
}

type HTTPConfig struct {
	ReadTimeout    time.Duration `mapstructure:"read_timeout" yaml:"read_timeout"`
	WriteTimeout   time.Duration `mapstructure:"write_timeout" yaml:"write_timeout"`
	IdleTimeout    time.Duration `mapstructure:"idle_timeout" yaml:"idle_timeout"`
	RequestTimeout time.Duration `mapstructure:"request_timeout" yaml:"request_timeout"`
}

type Logger struct {
	App          string `mapstructure:"app" yaml:"app"`
	Level        string `mapstructure:"level" yaml:"level"`
	Service      string `mapstructure:"service" yaml:"service"`
	UDPAddress   string `mapstructure:"udp_address" yaml:"udp_address"`
	EnableCaller bool   `mapstructure:"enable_caller" yaml:"enable_caller"`
}
