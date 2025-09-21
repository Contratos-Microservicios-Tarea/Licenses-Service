package env // Cambiar de "configs" a "env"

import (
	"os"
	"strconv"
	"time"

	logger "license-service/pkg/log/logger"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig `json:"database"`
	Server   ServerConfig   `json:"server"`
	App      AppConfig      `json:"app"`
}

type DatabaseConfig struct {
	Host            string        `json:"host"`
	Port            string        `json:"port"`
	Username        string        `json:"username"`
	Password        string        `json:"password"`
	Name            string        `json:"name"`
	SSLMode         string        `json:"ssl_mode"`
	TimeZone        string        `json:"timezone"`
	Url             string        `json:"url"`
	MaxIdleConns    int           `json:"max_idle_conns"`
	MaxOpenConns    int           `json:"max_open_conns"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`
	GormLogLevel    string        `json:"gorm_log_level"`
	SlowThreshold   time.Duration `json:"slow_threshold"`
}

type ServerConfig struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

type AppConfig struct {
	Environment string `json:"environment"`
	LogLevel    string `json:"log_level"`
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log := logger.NewLogger()
		log.Info("Environment Config", "Load", "No .env file found, using environment variables")
	}

	maxIdleConns := getEnvAsInt("DB_MAX_IDLE_CONNS", 10)
	maxOpenConns := getEnvAsInt("DB_MAX_OPEN_CONNS", 100)
	connMaxLifetime := time.Duration(getEnvAsInt("DB_CONN_MAX_LIFETIME", 3600)) * time.Second
	slowThreshold := time.Duration(getEnvAsInt("GORM_SLOW_THRESHOLD", 200)) * time.Millisecond

	return &Config{
		Database: DatabaseConfig{
			Host:            getEnv("POSTGRES_HOST", "localhost"),
			Port:            getEnv("POSTGRES_PORT", "5432"),
			Username:        getEnv("POSTGRES_USERNAME", "user"),
			Password:        getEnv("POSTGRES_PASSWORD", "password"),
			Name:            getEnv("POSTGRES_NAME", "health"),
			SSLMode:         getEnv("POSTGRES_SSLMODE", "disable"),
			TimeZone:        getEnv("POSTGRES_TIMEZONE", "America/Santiago"),
			Url:             getEnv("POSTGRES_URL", ""),
			MaxIdleConns:    maxIdleConns,
			MaxOpenConns:    maxOpenConns,
			ConnMaxLifetime: connMaxLifetime,
			GormLogLevel:    getEnv("GORM_LOG_LEVEL", "info"),
			SlowThreshold:   slowThreshold,
		},
		Server: ServerConfig{
			Port: getEnv("PORT", "8081"),
			Host: getEnv("HOST", "localhost"),
		},
		App: AppConfig{
			Environment: getEnv("APP_ENV", "development"),
			LogLevel:    getEnv("LOG_LEVEL", "info"),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
