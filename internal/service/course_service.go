package service

import (
	"app/internal/domain"
	"app/internal/repository"
)

type CourseService struct {
	courseRepo *repository.CourseRepository
}

func NewCourseService(courseRepo *repository.CourseRepository) *CourseService {
	return &CourseService{courseRepo: courseRepo}
}

func (s *CourseService) CreateCourse(course *domain.Course) error {
	// TODO: Extract from JWT context
	course.InstructorID = 1
	return s.courseRepo.Create(course)
}

func (s *CourseService) GetAllCourses() ([]domain.Course, error) {
	return s.courseRepo.GetAll()
}

func (s *CourseService) GetCourse(id int) (*domain.Course, error) {
	return s.courseRepo.GetByID(id)
}
