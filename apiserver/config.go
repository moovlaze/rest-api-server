package apiserver

type Config struct {
	Bindaddr string `yaml:"bind_addr"`
	LogLevel string `yaml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		Bindaddr: ":8080",
		LogLevel: "debug",
	}
}