package repository

import (
	"app/internal/domain"
	"app/internal/pkg/database"
	"database/sql"
)

type CourseRepository struct {
	DB *sql.DB
}

func (r *CourseRepository) Create(course *domain.Course) error {
	query := `INSERT INTO courses (title, description, instructor_id) VALUES (?, ?, ?)`
	result, err := database.DB.Exec(query, course.Title, course.Description, course.InstructorID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	course.ID = int(id)
	return nil
}

func (r *CourseRepository) GetAll() ([]domain.Course, error) {
	query := `SELECT id, title, description, instructor_id FROM courses`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []domain.Course
	for rows.Next() {
		var c domain.Course
		if err := rows.Scan(&c.ID, &c.Title, &c.Description, &c.InstructorID); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}

func (r *CourseRepository) GetByID(id int) (*domain.Course, error) {
	course := &domain.Course{}
	query := `SELECT id, title, description, instructor_id FROM courses WHERE id = ?`
	err := database.DB.QueryRow(query, id).Scan(&course.ID, &course.Title, &course.Description, &course.InstructorID)
	if err != nil {
		return nil, err
	}
	return course, nil
}
