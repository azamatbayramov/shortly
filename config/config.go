package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	AppHost string
	AppPort int

	StorageType string

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	CoderAlphabet string
	CoderLength   int

	OriginalLinkMaxLength int
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	cfg.AppHost = getEnv("APP_HOST", "0.0.0.0")
	cfg.AppPort = getEnvAsInt("APP_PORT", 8000)

	cfg.StorageType = getEnv("STORAGE_TYPE", "in_memory")

	cfg.PostgresHost = getEnv("POSTGRES_HOST", "")
	cfg.PostgresPort = getEnvAsInt("POSTGRES_PORT", 5432)
	cfg.PostgresDatabase = getEnv("POSTGRES_DATABASE", "")
	cfg.PostgresUser = getEnv("POSTGRES_USER", "")
	cfg.PostgresPassword = getEnv("POSTGRES_PASSWORD", "")

	cfg.CoderAlphabet = getEnv(
		"CODER_ALPHABET", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_",
	)
	cfg.CoderLength = getEnvAsInt("CODER_LENGTH", 10)

	cfg.OriginalLinkMaxLength = getEnvAsInt("ORIGINAL_LINK_MAX_LENGTH", 2048)

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		log.Printf("Invalid value for %s, using default: %d", key, defaultValue)
	}
	return defaultValue
}
