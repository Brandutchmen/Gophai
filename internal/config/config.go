package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv    string
	AppSecret string

	DbUrl string
}

var (
	config     *Config
	configOnce sync.Once
)

func Init() {
	if _, err := GetConfig(); err != nil {
		panic("Failed to load .env file")
	}
}

func GetConfig() (*Config, error) {
	var err error
	configOnce.Do(func() {
		config, err = loadEnvFile()
	})
	return config, err
}

func SetConfig(cfg Config) Config {
	config = &cfg
	return cfg
}

func loadEnvFile() (*Config, error) {

	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "local"
	}

	if env != "testing" {
		fmt.Println("Loading .env file from disk")
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("Error loading .env file: %w", err)
		}
	}

	config := &Config{
		AppEnv:    os.Getenv("APP_ENV"),
		AppSecret: os.Getenv("APP_SECRET"),
	}
	config.DbUrl = os.Getenv("DB_URL")

	return config, nil
}
