package storage

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/krJay1/go-helpdesk/internal/config"
)

// For golang-migrate
func dbURL(cfg *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
}
func MigrateDb(cfg *config.Config) error {
	m, err := migrate.New(
		"file://migrations",
		dbURL(cfg),
	)
	if err != nil {
		return err
	}

	defer func() {
		sourceErr, dbErr := m.Close()

		if sourceErr != nil {
			log.Printf("migration source close error: %v", sourceErr)
		}

		if dbErr != nil {
			log.Printf("migration db close error: %v", dbErr)
		}
	}()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Println("Migrations applied")
	return nil
}
