package config

type Config struct {
	RequestConfig RequestConfig
}

func NewConfig() *Config {
	return &Config{
		RequestConfig:RequestConfig{
			DefaultFormat:"application/x-www-form-urlencoded",
		},
	};
}
