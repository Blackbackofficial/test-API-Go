package apiserver

import "awesomeProject/internal/store"

type Config struct {
	StateAddr string `toml:"state_addr"`
	LogLevel string `toml:"log_level"`
	Store *store.Config
}

func NewConfig() *Config {
	return &Config{
		StateAddr: ":8080",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}