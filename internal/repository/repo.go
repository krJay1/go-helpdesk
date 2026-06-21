package repository

import "database/sql"

type AppRepository struct {
	DB *sql.DB
}
