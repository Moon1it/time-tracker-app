package database

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(connString string) error {
	m, err := migrate.New(
		"file://internal/db/migrations",
		// "postgres://user:password@localhost:5432/time_tracker_app_db?sslmode=disable")
		connString)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
