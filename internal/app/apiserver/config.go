package apiserver

type Config struct {
	StateAddr string `toml:"state_addr"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		StateAddr: ":8080",
		LogLevel: "debug",
	}
}