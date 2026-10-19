package config

import "testing"

func clearConfigEnv(t *testing.T) {
	t.Helper()

	for _, key := range []string{
		"SERVER_HOST",
		"SERVER_PORT",
		"MAX_REQUEST_BYTES",
		"MAX_EXPRESSION_LENGTH",
		"HTTP_READ_TIMEOUT",
		"HTTP_WRITE_TIMEOUT",
		"HTTP_IDLE_TIMEOUT",
	} {
		t.Setenv(key, "")
	}
}

func TestDefaultConfig(t *testing.T) {
	clearConfigEnv(t)

	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}

	if cfg.ServerHost != "0.0.0.0" {
		t.Fatalf("unexpected host: %s", cfg.ServerHost)
	}

	if cfg.ServerPort != 8080 {
		t.Fatalf("unexpected port: %d", cfg.ServerPort)
	}

	if cfg.MaxRequestBytes != 4096 {
		t.Fatalf(
			"unexpected request limit: %d",
			cfg.MaxRequestBytes,
		)
	}

	if cfg.MaxExpressionLength != 1024 {
		t.Fatalf(
			"unexpected expression limit: %d",
			cfg.MaxExpressionLength,
		)
	}
}

func TestCustomConfig(t *testing.T) {
	clearConfigEnv(t)

	t.Setenv("SERVER_PORT", "9090")
	t.Setenv("MAX_REQUEST_BYTES", "8192")
	t.Setenv("MAX_EXPRESSION_LENGTH", "2048")

	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}

	if cfg.ServerPort != 9090 {
		t.Fatalf("unexpected port: %d", cfg.ServerPort)
	}

	if cfg.MaxRequestBytes != 8192 {
		t.Fatalf(
			"unexpected request limit: %d",
			cfg.MaxRequestBytes,
		)
	}

	if cfg.MaxExpressionLength != 2048 {
		t.Fatalf(
			"unexpected expression limit: %d",
			cfg.MaxExpressionLength,
		)
	}
}

func TestInvalidPort(t *testing.T) {
	clearConfigEnv(t)
	t.Setenv("SERVER_PORT", "99999")

	_, err := Load()
	if err == nil {
		t.Fatal("expected invalid port error")
	}
}

func TestInvalidTimeout(t *testing.T) {
	clearConfigEnv(t)
	t.Setenv("HTTP_READ_TIMEOUT", "-5s")

	_, err := Load()
	if err == nil {
		t.Fatal("expected invalid timeout error")
	}
}
