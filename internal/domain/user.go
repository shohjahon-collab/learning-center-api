package domain

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	Phone        string `json:"phone"`
	PasswordHash string `json:"-"` // not exposed
	Role         string `json:"role"`
}
