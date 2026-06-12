package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/krJay1/go-helpdesk/internal/config"
)

var DB *sql.DB

func InitDB(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	fmt.Println("ConnectionString:", connStr)

	DB, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}
	if err := DB.Ping(); err != nil {
		return nil, err
	}

	return DB, nil

}
