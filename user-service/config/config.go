package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL      string
	ServerPort string
	JWTSecret  string
}

func Load() *Config {
	_ = godotenv.Load()
	return &Config{
		DBURL:      os.Getenv("DATABASE_URL"),
		ServerPort: os.Getenv("SERVER_PORT"),
		JWTSecret:  os.Getenv("SECRET_KEY"),
	}
}
