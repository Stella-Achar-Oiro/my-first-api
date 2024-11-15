package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	ServerPort     string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	AllowedOrigins []string
	LogLevel       string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort:     getEnv("SERVER_PORT", "8080"),
		ReadTimeout:    time.Duration(getEnvAsInt("READ_TIMEOUT", 10)) * time.Second,
		WriteTimeout:   time.Duration(getEnvAsInt("WRITE_TIMEOUT", 10)) * time.Second,
		AllowedOrigins: getEnvAsSlice("ALLOWED_ORIGINS", []string{"*"}),
		LogLevel:       getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string) []string {
	if value, exists := os.LookupEnv(key); exists {
		return strings.Split(value, ",")
	}
	return defaultValue
}
