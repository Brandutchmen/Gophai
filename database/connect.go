package database

import (
	"app/config"
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("failed to parse DB_PORT to int")
	}

	dsn := fmt.Sprintf(
		"host=db port=%d user=%s password=%s dbname=%s sslmode=disable",
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}

	fmt.Println("Connected to the database")
	DB.AutoMigrate()
}
