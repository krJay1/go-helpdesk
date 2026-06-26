package repository

import "github.com/jackc/pgx/v5/pgxpool"

type AppRepository struct {
	DB *pgxpool.Pool
}
