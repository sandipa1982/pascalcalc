package config

import "testing"

func TestLoadForService(t *testing.T) {
	t.Setenv("SERVER_PORT", "")
	cfg, err := LoadForService(8081)
	if err != nil || cfg.ServerPort != 8081 {
		t.Fatalf("default port: %+v %v", cfg, err)
	}
	t.Setenv("SERVER_PORT", "9191")
	cfg, err = LoadForService(8081)
	if err != nil || cfg.ServerPort != 9191 {
		t.Fatalf("override port: %+v %v", cfg, err)
	}
}
