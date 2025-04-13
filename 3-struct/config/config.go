package config

import "os"

type Config struct {
	ApiKey string
}

func NewConfig() *Config {
	apiKey := os.Getenv("APIKEY")
	if apiKey == "" {
		panic("APIKEY environment variable is not set")
	}
	return &Config{ApiKey: apiKey}
}
