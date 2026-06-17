package models

type User struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	MobileNumber string `json:"mobile_number"`
	PasswordHash string `json:"-"`
	LastLogin    string `json:"last_login"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	IsActive     bool   `json:"is_active"`
}
