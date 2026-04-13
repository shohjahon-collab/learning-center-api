package repository

import (
	"database/sql"
	"learning-center-api/internal/domain"
	"learning-center-api/internal/pkg/database"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) Create(user *domain.User) error {
	query := `INSERT INTO users (email, full_name, phone, password_hash, role) VALUES (?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(query, user.Email, user.FullName, user.Phone, user.PasswordHash, user.Role)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, email, full_name, phone, password_hash, role FROM users WHERE email = ?`
	err := database.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.FullName, &user.Phone, &user.PasswordHash, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}
