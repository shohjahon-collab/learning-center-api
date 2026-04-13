package config

import (
	"os"
)

type Config struct {
	Port      string
	JWTSecret string
	DBPath    string
}

func Load() *Config {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "my-super-secret-jwt-key-change-in-prod"
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "learning_center.db"
	}

	return &Config{
		Port:      portStr,
		JWTSecret: jwtSecret,
		DBPath:    dbPath,
	}
}
