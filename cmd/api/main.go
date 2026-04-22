package main

import (
	"log"
	"net/http"

	"app/internal/config"
	"app/internal/handler"
	"app/internal/middleware"
	"app/internal/pkg/database"
	"app/internal/repository"
	"app/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Load()

	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}
	defer database.DB.Close()

	userRepo := &repository.UserRepository{}
	courseRepo := &repository.CourseRepository{}

	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	courseService := service.NewCourseService(courseRepo)

	authHandler := handler.NewAuthHandler(authService)
	courseHandler := handler.NewCourseHandler(courseService)

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	api.Handle("/auth/login", middleware.CORS(authHandler.Login())).Methods("POST")
	api.Handle("/auth/register", middleware.CORS(authHandler.Register())).Methods("POST")

	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.Authenticate(cfg.JWTSecret))

	protected.Handle("/courses", middleware.CORS(courseHandler.CreateCourse())).Methods("POST")
	protected.Handle("/courses", middleware.CORS(courseHandler.GetAllCourses())).Methods("GET")
	protected.Handle("/courses/{id}", middleware.CORS(courseHandler.GetCourse())).Methods("GET")

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
