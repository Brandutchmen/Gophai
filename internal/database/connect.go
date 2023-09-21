package database

import (
	"app/internal/config"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connect() {
	var err error

	config, err := config.GetConfig()
	if err != nil {
		panic("Failed to load config when connecting to the database")
	}

	conn, err := pgx.Connect(context.Background(), config.DbUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Connected to the database")
}
