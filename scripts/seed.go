package main

import (
	"log"

	"learning-center-api/internal/domain"
	"learning-center-api/internal/pkg/database"
	"learning-center-api/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	database.InitDB()

	userRepo := &repository.UserRepository{}

	// Create admin user
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Admin123!"), bcrypt.DefaultCost)

	admin := &domain.User{
		Email:        "admin@learningcenter.com",
		PasswordHash: string(hashedPassword),
		FullName:     "System Administrator",
		Role:         "admin",
	}

	if err := userRepo.Create(admin); err != nil {
		log.Println("Admin already exists or error:", err)
	} else {
		log.Println("Admin user created successfully!")
		log.Println("Email: admin@learningcenter.com")
		log.Println("Password: Admin123!")
	}
}
