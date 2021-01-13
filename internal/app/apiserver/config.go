package apiserver

type Config struct {
	StateAddr string `toml:"state_addr"`
}

func NewConfig() *Config {
	return &Config{
		StateAddr: ":8080",
	}
}