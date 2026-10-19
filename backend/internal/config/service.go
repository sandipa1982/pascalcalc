package config

import "os"

// LoadForService applies a service-specific default port while preserving
// SERVER_PORT overrides and all validation performed by Load.
func LoadForService(defaultPort int) (Config, error) {
	cfg, err := Load()
	if err != nil {
		return Config{}, err
	}
	if os.Getenv("SERVER_PORT") == "" {
		cfg.ServerPort = defaultPort
	}
	return cfg, nil
}
