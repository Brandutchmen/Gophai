package config_test

import (
	"testing"

	"app/internal/config"
)

func setupTestEnv() config.Config {
	return config.SetConfig(config.Config{
		AppEnv:    "production",
		AppSecret: "my-secret",
		DbUrl:     "postgres://username:password@localhost:5432/database_name",
	})
}

func TestGetConfig(t *testing.T) {

	cfg := setupTestEnv()

	// Check the config values
	if cfg.AppEnv != "production" {
		t.Errorf("Expected AppEnv to be 'production', got %s", cfg.AppEnv)
	}
	if cfg.AppSecret != "my-secret" {
		t.Errorf("Expected AppSecret to be 'my-secret', got %s", cfg.AppSecret)
	}
	if cfg.DbUrl != "postgres://username:password@localhost:5432/database_name" {
		t.Errorf("Expected DbUrl to be 'postgres://username:password@localhost:5432/database_name', got %s", cfg.DbUrl)
	}
}

func TestConfigIsSingleton(t *testing.T) {
	// Get the config
	cfg, err := config.GetConfig()
	if err != nil {
		t.Errorf("Error getting config: %v", err)
	}

	// Check that the config is a singleton
	cfg2, err := config.GetConfig()
	if err != nil {
		t.Errorf("Error getting config: %v", err)
	}
	if cfg != cfg2 {
		t.Errorf("Expected config to be a singleton")
	}
}
