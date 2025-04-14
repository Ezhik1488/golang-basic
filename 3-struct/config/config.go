package config

import "os"

type Config struct {
	ApiKey string
	ApiUrl string
}

func NewConfig() *Config {
	apiKey := os.Getenv("API_KEY")
	apiUrl := os.Getenv("API_URL")
	if apiKey == "" || apiUrl == "" {
		panic("APIKEY || APIURL environment variable is not set")
	}
	return &Config{ApiKey: apiKey, ApiUrl: apiUrl}
}
