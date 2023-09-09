package database

import (
	"app/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	var err error

	config, err := config.GetConfig()
	if err != nil {
		panic("Failed to load config when connecting to the database")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPassword,
		config.DbName,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}

	fmt.Println("Connected to the database")
	DB.AutoMigrate()
}
