package config

import (
	"os"
)

type Config struct {
	DBDSN      string
	ServerPort string
}

func Load() *Config {
	return &Config{
		DBDSN:      os.Getenv("DB_DSN"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
