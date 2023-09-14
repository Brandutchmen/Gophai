package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv    string
	AppSecret string

	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
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

	// Parse integer environment variables
	if portStr := os.Getenv("DB_PORT"); portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("Error parsing DB_PORT: %w", err)
		}
		config.DbPort = port
	}

	config.DbHost = os.Getenv("DB_HOST")
	config.DbUser = os.Getenv("DB_USER")
	config.DbPassword = os.Getenv("DB_PASSWORD")
	config.DbName = os.Getenv("DB_NAME")

	return config, nil
}
