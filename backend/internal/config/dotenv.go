package config

import (
	"errors"
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

// Existing OS environment variables take precedence.
func LoadDotEnv() error {
	path := os.Getenv("ENV_FILE")

	if path == "" {
		path = "../.env"
	}

	err := godotenv.Load(path)
	if err == nil {
		return nil
	}

	if errors.Is(err, os.ErrNotExist) {
		// The file is optional when environment variables
		// are provided by Docker or the operating system.
		return nil
	}

	return fmt.Errorf("load environment file %q: %w", path, err)
}
