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
		`INSERT INTO users(first_name, last_name, email) 
		VALUES($1, $2, $3)
		RETURNING id`,
		user.FirstName,
		user.LastName,
		user.Email,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (r *UserRepository) GetUser(id int64) (*models.User, error) {
	user := &models.User{}

	err := r.DB.QueryRow(
		"SELECT id, first_name, last_name, email, mobile_number, last_login, created_at, updated_at, is_active, password_hash FROM users WHERE id=$1",
		id,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.MobileNumber,
		&user.LastLogin,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsActive,
		&user.PasswordHash,
	)

	if err != nil {
		return nil, err
	}
	return user, err
}
