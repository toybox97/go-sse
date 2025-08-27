package config

import (
	"os"
)

type Config struct {
	Port string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	// めんどくさいので環境変数忘れた時は直入れする
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port: port,
	}
}

func (c *Config) GetAddr() string {
	return ":" + c.Port
}
