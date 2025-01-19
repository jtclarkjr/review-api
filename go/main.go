package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swaggo/http-swagger"
	"go-api/db"
	_ "go-api/docs"
	"go-api/handlers"
	"go-api/middlewares"
)

// @title Go API
// @version 1.0
// @description This is a sample server for a GO API.

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Provide your token with prefix "Bearer "
// @schemes https http
func main() {
	// Initialize database connection
	db.Connect()
	db.SeedDatabase()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Swagger route
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// Login to generate token route
	r.Post("/login", handlers.Login)

	r.Route("/admin", func(r chi.Router) {
		r.Use(middlewares.AuthAdmin)
		r.Post("/employees", handlers.AddEmployee)
		r.Get("/employees", handlers.GetEmployees)
		r.Put("/employees/{id}", handlers.UpdateEmployee)
		r.Delete("/employees/{id}", handlers.RemoveEmployee)

		r.Post("/reviews", handlers.AddReview)
		r.Get("/reviews", handlers.GetReviews)
		r.Put("/reviews/{id}/comments", handlers.UpdateReview)
	})

	r.Route("/employee", func(r chi.Router) {
		r.Use(middlewares.AuthEmployee)
		r.Get("/reviews", handlers.ListReviews)
		r.Post("/reviews/feedback", handlers.SubmitFeedback)
	})

	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
