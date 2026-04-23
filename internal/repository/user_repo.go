package repository

import (
	"app/internal/domain"
	"app/internal/pkg/database"
)

type UserRepository struct{}

func (r *UserRepository) Create(user *domain.User) error {
	query := `INSERT INTO users (email, full_name, phone, password_hash, role) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := database.DB.QueryRow(query, user.Email, user.FullName, user.Phone, user.PasswordHash, user.Role).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, email, full_name, phone, password_hash, role FROM users WHERE email = $1`
	err := database.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.FullName, &user.Phone, &user.PasswordHash, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}
