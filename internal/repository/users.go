package repository

import (
	"context"

	"github.com/krJay1/go-helpdesk/internal/models"
)

func (r *AppRepository) CreateUser(ctx context.Context, user models.User) (int64, error) {
	var id int64
	err := r.DB.QueryRow(
		ctx,
		`INSERT INTO users(first_name, last_name, email, mobile_number, password_hash) 
		VALUES($1, $2, $3, $4, $5)
		RETURNING id`,
		user.FirstName,
		user.LastName,
		user.Email,
		user.MobileNumber,
		user.PasswordHash,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (r *AppRepository) getUser(ctx context.Context, query string, i interface{}) (models.User, error) {
	var user models.User

	err := r.DB.QueryRow(
		ctx,
		query,
		i,
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
		return models.User{}, err
	}
	return user, nil
}

func (r *AppRepository) GetUsers(ctx context.Context) ([]models.User, error) {
	rows, err := r.DB.Query(
		ctx,
		`SELECT 
		id, 
		first_name, 
		last_name, 
		email, 
		mobile_number, 
		last_login, 
		created_at, 
		updated_at, 
		is_active FROM users;`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.MobileNumber,
			&user.LastLogin,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.IsActive,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *AppRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	query := "SELECT id, first_name, last_name, email, mobile_number, last_login, created_at, updated_at, is_active, password_hash FROM users WHERE email=$1;"

	user, err := r.getUser(ctx, query, email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *AppRepository) GetUserByID(ctx context.Context, id int64) (models.User, error) {
	var user models.User

	query := "SELECT id, first_name, last_name, email, mobile_number, last_login, created_at, updated_at, is_active, password_hash FROM users WHERE id=$1;"
	user, err := r.getUser(ctx, query, id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
