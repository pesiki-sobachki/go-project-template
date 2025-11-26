package config

type Config struct {
	Env    string `mapstructure:"env"`
	Logger Logger `mapstructure:"logger"`
}

type Logger struct {
	App          string `mapstructure:"app"`
	Level        string `mapstructure:"level"`
	Service      string `mapstructure:"service"`
	UDPAddress   string `mapstructure:"udp_address"`
	EnableCaller bool   `mapstructure:"enable_caller"`
}
