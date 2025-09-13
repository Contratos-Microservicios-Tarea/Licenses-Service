package configs

import (
	"os"

	logger "license-service/pkg/log/logger"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig `json:"database"`
	Server   ServerConfig   `json:"server"`
}

type DatabaseConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Url      string `json:"url"`
	Name     string `json:"name"`
}

type ServerConfig struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		logger.Warn(
			"Environment Config",
			"Loading .env file",
			"No .env file found, using environment variables",
		)
	}

	return &Config{
		Database: DatabaseConfig{
			Username: getEnv("POSTGRES_USERNAME", ""),
			Password: getEnv("POSTGRES_PASSWORD", ""),
			Url:      getEnv("POSTGRES_URL", ""),
			Name:     getEnv("POSTGRES_NAME", ""),
		},
		Server: ServerConfig{
			Port: getEnv("PORT", "8081"),
			Host: getEnv("HOST", "localhost"),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
