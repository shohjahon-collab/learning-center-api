package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	userTable := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email TEXT UNIQUE NOT NULL,
			full_name TEXT NOT NULL,
			phone TEXT,
			password_hash TEXT NOT NULL,
			role TEXT DEFAULT 'student' CHECK(role IN ('student', 'instructor', 'admin'))
		);
	`
	courseTable := `
		CREATE TABLE IF NOT EXISTS courses (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT,
			instructor_id INTEGER,
			FOREIGN KEY (instructor_id) REFERENCES users (id)
		);
	`

	_, err = DB.Exec(userTable)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}
	_, err = DB.Exec(courseTable)
	if err != nil {
		return fmt.Errorf("failed to create courses table: %w", err)
	}

	log.Println("Database initialized successfully")
	return nil
}
