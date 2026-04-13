package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"learning-center-api/internal/domain"
	"learning-center-api/internal/service"

	"github.com/gorilla/mux"
)

type CourseHandler struct {
	courseService *service.CourseService
}

func NewCourseHandler(courseService *service.CourseService) *CourseHandler {
	return &CourseHandler{courseService: courseService}
}

type CreateCourseRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *CourseHandler) CreateCourse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateCourseRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		course := &domain.Course{
			Title:       req.Title,
			Description: req.Description,
		}
		if err := h.courseService.CreateCourse(course); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(course)
	}
}

func (h *CourseHandler) GetAllCourses() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courses, err := h.courseService.GetAllCourses()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(courses)
	}
}

func (h *CourseHandler) GetCourse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid course ID", http.StatusBadRequest)
			return
		}

		course, err := h.courseService.GetCourse(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(course)
	}
}
