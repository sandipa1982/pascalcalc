package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config contains application configuration.
// Values can be overridden using environment variables.
type Config struct {
	ServerHost          string
	ServerPort          int
	MaxRequestBytes     int64
	MaxExpressionLength int
	ReadTimeout         time.Duration
	WriteTimeout        time.Duration
	IdleTimeout         time.Duration
}

// Load reads environment variables, applies defaults,
// and validates the configuration.
func Load() (Config, error) {
	cfg := Config{
		ServerHost:          getEnv("SERVER_HOST", "0.0.0.0"),
		ServerPort:          8080,
		MaxRequestBytes:     4096,
		MaxExpressionLength: 1024,
		ReadTimeout:         5 * time.Second,
		WriteTimeout:        10 * time.Second,
		IdleTimeout:         60 * time.Second,
	}

	var err error

	cfg.ServerPort, err = getEnvInt(
		"SERVER_PORT", cfg.ServerPort,
	)
	if err != nil {
		return Config{}, err
	}

	cfg.MaxRequestBytes, err = getEnvInt64(
		"MAX_REQUEST_BYTES", cfg.MaxRequestBytes,
	)
	if err != nil {
		return Config{}, err
	}

	cfg.MaxExpressionLength, err = getEnvInt(
		"MAX_EXPRESSION_LENGTH", cfg.MaxExpressionLength,
	)
	if err != nil {
		return Config{}, err
	}

	cfg.ReadTimeout, err = getEnvDuration(
		"HTTP_READ_TIMEOUT", cfg.ReadTimeout,
	)
	if err != nil {
		return Config{}, err
	}

	cfg.WriteTimeout, err = getEnvDuration(
		"HTTP_WRITE_TIMEOUT", cfg.WriteTimeout,
	)
	if err != nil {
		return Config{}, err
	}

	cfg.IdleTimeout, err = getEnvDuration(
		"HTTP_IDLE_TIMEOUT", cfg.IdleTimeout,
	)
	if err != nil {
		return Config{}, err
	}

	if cfg.ServerPort < 1 || cfg.ServerPort > 65535 {
		return Config{}, fmt.Errorf(
			"SERVER_PORT must be between 1 and 65535",
		)
	}

	if cfg.ServerHost == "" {
		return Config{}, fmt.Errorf("SERVER_HOST cannot be empty")
	}

	if cfg.MaxRequestBytes <= 0 ||
		cfg.MaxExpressionLength <= 0 {
		return Config{}, fmt.Errorf(
			"request and expression limits must be positive",
		)
	}

	if cfg.ReadTimeout <= 0 ||
		cfg.WriteTimeout <= 0 ||
		cfg.IdleTimeout <= 0 {
		return Config{}, fmt.Errorf(
			"HTTP timeouts must be positive",
		)
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getEnvInt(key string, fallback int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", key, err)
	}
	return parsed, nil
}

func getEnvInt64(key string, fallback int64) (int64, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", key, err)
	}
	return parsed, nil
}

func getEnvDuration(
	key string,
	fallback time.Duration,
) (time.Duration, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	parsed, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", key, err)
	}
	return parsed, nil
}
