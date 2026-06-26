package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/krJay1/go-helpdesk/internal/config"
)

var DB *pgxpool.Pool

func InitDB(cfg *config.Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	fmt.Println("ConnectionString:", connStr)

	conf, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	conf.MaxConns = 10
	conf.MinConns = 2

	conf.HealthCheckPeriod = time.Minute // 30 seconds
	conf.MaxConnLifetime = 30 * time.Minute
	conf.MaxConnIdleTime = 15 * time.Minute
	// conf.ConnConfig.ConnectTimeout = 5 * time.Second

	DB, err = pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		return nil, err
	}

	if err := DB.Ping(context.Background()); err != nil {
		return nil, err
	}

	return DB, nil

}
