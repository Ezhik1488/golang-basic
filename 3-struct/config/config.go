package config

import "os"

type Config struct {
	ApiKey           string
	ApiUrl           string
	LocalStoragePath string
}

func NewConfig() *Config {
	apiKey := os.Getenv("API_KEY")
	apiUrl := os.Getenv("API_URL")
	localStoragePath := os.Getenv("LOCAL_STORAGE_PATH")
	if apiKey == "" || apiUrl == "" {
		panic("APIKEY || APIURL environment variable is not set")
	}
	return &Config{ApiKey: apiKey, ApiUrl: apiUrl, LocalStoragePath: localStoragePath}
}
