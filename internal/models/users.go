package models

import "time"

type User struct {
	ID           int64      `json:"id"`
	Email        string     `json:"email"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	MobileNumber *string    `json:"mobile_number"`
	PasswordHash string     `json:"-"`
	LastLogin    *time.Time `json:"last_login"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	IsActive     bool       `json:"is_active"`
}
