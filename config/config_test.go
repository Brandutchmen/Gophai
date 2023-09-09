package config_test

import (
	"testing"

	"app/config"
)

func setupTestEnv() {
	config.GetConfig()
	// Set up the configuration for testing
	config.SetConfig(config.Config{
		AppEnv:     "production",
		AppSecret:  "my-secret",
		DbPort:     5432,
		DbHost:     "localhost",
		DbUser:     "postgres",
		DbPassword: "my-password",
		DbName:     "my-db",
	})
}

func TestGetConfig(t *testing.T) {

	setupTestEnv()

	// Get the config
	cfg, err := config.GetConfig()
	if err != nil {
		t.Errorf("Error getting config: %v", err)
	}

	// Check the config values
	if cfg.AppEnv != "production" {
		t.Errorf("Expected AppEnv to be 'production', got %s", cfg.AppEnv)
	}
	if cfg.AppSecret != "my-secret" {
		t.Errorf("Expected AppSecret to be 'my-secret', got %s", cfg.AppSecret)
	}
	if cfg.DbPort != 5432 {
		t.Errorf("Expected DbPort to be 5432, got %d", cfg.DbPort)
	}
	if cfg.DbHost != "localhost" {
		t.Errorf("Expected DbHost to be 'localhost', got %s", cfg.DbHost)
	}
	if cfg.DbUser != "postgres" {
		t.Errorf("Expected DbUser to be 'postgres', got %s", cfg.DbUser)
	}
	if cfg.DbPassword != "my-password" {
		t.Errorf("Expected DbPassword to be 'my-password', got %s", cfg.DbPassword)
	}
	if cfg.DbName != "my-db" {
		t.Errorf("Expected DbName to be 'my-db', got %s", cfg.DbName)
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
