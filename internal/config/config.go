package config

type Config struct {
	Env    string `mapstructure:"env" yaml:"env"`
	Addr   string `mapstructure:"addr" yaml:"addr"`
	Logger Logger `mapstructure:"logger" yaml:"logger"`
}

type Logger struct {
	App          string `mapstructure:"app" yaml:"app"`
	Level        string `mapstructure:"level" yaml:"level"`
	Service      string `mapstructure:"service" yaml:"service"`
	UDPAddress   string `mapstructure:"udp_address" yaml:"udp_address"`
	EnableCaller bool   `mapstructure:"enable_caller" yaml:"enable_caller"`
}
