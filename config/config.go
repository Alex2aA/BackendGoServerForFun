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
		DBURL:      getEnv("DATABASE_URL", "postgres://postgres:password@localhost:8348/golang_web_server"),
		ServerPort: getEnv("SERVER_PORT", ":8080"),
		JWTSecret:  getEnv("SECRET_KEY", "JWT_SECRET"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
