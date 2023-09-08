package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}

func ConfigWithDefault(key string, defaultValue string) string {
	value := Config(key)

	if value == "" {
		value = defaultValue
	}

	return value
}
