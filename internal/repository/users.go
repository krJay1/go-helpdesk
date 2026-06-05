package repository

import (
	"database/sql"

	"github.com/krJay1/go-helpdesk/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user models.User) (int64, error) {
	var id int64
	err := r.DB.QueryRow(
		`INSERT INTO users(name, email) 
		VALUES($1, $2)
		RETURNING id`,
		user.Name,
		user.Email,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (r *UserRepository) GetUser(id int) (*models.User, error) {
	user := &models.User{}

	err := r.DB.QueryRow(
		"SELECT id, name, email FROM users WHERE id=$1",
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	if err != nil {
		return nil, err
	}
	return user, err
}
