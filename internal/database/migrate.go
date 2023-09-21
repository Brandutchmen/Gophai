package database

import (
	"app/internal/config"
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate() {
	m, err := GetMigrateInstance()
	if err != nil {
		panic(err)
	}
	m.Up()
}

func TestLastMigration() {
	m, err := GetMigrateInstance()
	if err != nil {
		panic(err)
	}
	m.Steps(1)
	m.Steps(-1)
	m.Steps(1)
}

func GetMigrateInstance() (*migrate.Migrate, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/database/migrations",
		"postgres", driver)
	return m, err
}
