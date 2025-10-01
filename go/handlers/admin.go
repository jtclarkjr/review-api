package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"go-api/db"
	"go-api/types"

	"github.com/jtclarkjr/router-go"
	"github.com/lib/pq"
)

// /employees handlers

// AddEmployee godoc
// @Summary Add a new employee
// @Description Adds a new employee to the system
// @Tags Admin
// @Accept json
// @Produce json
// @Param employee body object true "Employee info"
// @Success 201 {object} types.CreateEmployeeResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/employees [post]
func AddEmployee(w http.ResponseWriter, r *http.Request) {
	var employee struct {
		Email    string `json:"email"`
		Position string `json:"position"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := hashPassword(employee.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	tx, err := db.Conn.Begin()
	if err != nil {
		http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return
	}

	var employeeID int
	err = tx.QueryRow(
		"INSERT INTO employees (email, position) VALUES ($1, $2) RETURNING id",
		employee.Email, employee.Position,
	).Scan(&employeeID)
	if err != nil {
		_ = tx.Rollback()
		http.Error(w, "Error adding employee", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec(
		"INSERT INTO users (email, password, role) VALUES ($1, $2, 'employee')",
		employee.Email, hashedPassword,
	)
	if err != nil {
		_ = tx.Rollback()
		http.Error(w, "Error creating user account", http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, "Error committing transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(types.CreateEmployeeResponse{
		EmployeeID: employeeID,
		Email:      employee.Email,
	})
	if err != nil {
		return
	}
}

// GetEmployees godoc
// @Summary Get all employees
// @Description Retrieves a list of all employees
// @Tags Admin
// @Produce json
// @Success 200 {array} types.EmployeeResponse
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/employees [get]
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Conn.Query("SELECT id, email, position FROM employees")
	if err != nil {
		http.Error(w, "Error fetching employees", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

	var employees []types.EmployeeResponse
	for rows.Next() {
		var id int
		var email, position string
		err := rows.Scan(&id, &email, &position)
		if err != nil {
			http.Error(w, "Error scanning employee data", http.StatusInternalServerError)
			return
		}
		employees = append(employees, types.EmployeeResponse{
			ID:       id,
			Email:    email,
			Position: position,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(employees)
	if err != nil {
		return
	}
}

// UpdateEmployee godoc
// @Summary Update an employee
// @Description Updates an employee's information
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body object true "Employee info"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/employees/{id} [put]
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	employeeID := router.URLParam(r, "id")

	var employee struct {
		Email    string `json:"email"`
		Position string `json:"position"`
	}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, err = db.Conn.Exec(
		"UPDATE employees SET email = $1, position = $2 WHERE id = $3",
		employee.Email, employee.Position, employeeID,
	)
	if err != nil {
		http.Error(w, "Error updating employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RemoveEmployee godoc
// @Summary Remove an employee
// @Description Removes an employee from the system
// @Tags Admin
// @Param id path int true "Employee ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/employees/{id} [delete]
func RemoveEmployee(w http.ResponseWriter, r *http.Request) {
	employeeID := router.URLParam(r, "id")

	_, err := db.Conn.Exec("DELETE FROM employees WHERE id = $1", employeeID)
	if err != nil {
		http.Error(w, "Error removing employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// /review handlers

// AddReview godoc
// @Summary Add a new review
// @Description Creates a new performance review and assigns reviewers
// @Tags Admin
// @Accept json
// @Produce json
// @Param review body object true "Review info"
// @Success 201 {object} types.CreateReviewResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/reviews [post]
func AddReview(w http.ResponseWriter, r *http.Request) {
	var review struct {
		EmployeeID        int    `json:"employee_id"`        // Employee being reviewed
		PerformanceReview string `json:"performance_review"` // Review text
		ReviewerIDs       []int  `json:"reviewer_ids"`       // List of reviewers
	}
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	tx, err := db.Conn.Begin()
	if err != nil {
		http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return
	}

	// Insert the review into the database
	var reviewID int
	err = tx.QueryRow(
		"INSERT INTO reviews (employee_id, performance_review, comments) VALUES ($1, $2, $3) RETURNING id",
		review.EmployeeID, review.PerformanceReview, pq.Array([]string{}),
	).Scan(&reviewID)
	if err != nil {
		_ = tx.Rollback()
		http.Error(w, "Error adding review", http.StatusInternalServerError)
		return
	}

	// Add reviewers to the review_reviewers table concurrently
	errChan := make(chan error, len(review.ReviewerIDs)) // Buffered channel for errors
	var wg sync.WaitGroup

	for _, reviewerID := range review.ReviewerIDs {
		wg.Go(func() {
			_, err := tx.Exec(
				"INSERT INTO review_reviewers (review_id, reviewer_id) VALUES ($1, $2)",
				reviewID, reviewerID,
			)
			errChan <- err
		})
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(errChan)

	// Check for errors from the goroutines
	for err := range errChan {
		if err != nil {
			_ = tx.Rollback()
			http.Error(w, "Error adding reviewers", http.StatusInternalServerError)
			return
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		http.Error(w, "Error committing transaction", http.StatusInternalServerError)
		return
	}

	// Respond with the review ID
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(types.CreateReviewResponse{
		ReviewID: reviewID,
	})
	if err != nil {
		http.Error(w, "Error encoding response for add review", http.StatusInternalServerError)
	}
}

// UpdateReview godoc
// @Summary Update a review
// @Description Updates the performance review text and reviewers
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path int true "Review ID"
// @Param review body object true "Review info"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/reviews/{id}/comments [put]
func UpdateReview(w http.ResponseWriter, r *http.Request) {
	reviewID := router.URLParam(r, "id")

	var payload struct {
		PerformanceReview string `json:"performance_review"` // Updated review text
		ReviewerIDs       []int  `json:"reviewer_ids"`       // List of new reviewers
	}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	tx, err := db.Conn.Begin()
	if err != nil {
		http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return
	}

	// Update the performance review
	_, err = tx.Exec(
		"UPDATE reviews SET performance_review = $1 WHERE id = $2",
		payload.PerformanceReview, reviewID,
	)
	if err != nil {
		_ = tx.Rollback()
		http.Error(w, "Error updating review", http.StatusInternalServerError)
		return
	}

	// Clear existing reviewers
	_, err = tx.Exec("DELETE FROM review_reviewers WHERE review_id = $1", reviewID)
	if err != nil {
		_ = tx.Rollback()
		http.Error(w, "Error clearing existing reviewers", http.StatusInternalServerError)
		return
	}

	// Add new reviewers concurrently
	errChan := make(chan error, len(payload.ReviewerIDs)) // Buffered channel for errors
	var wg sync.WaitGroup

	for _, reviewerID := range payload.ReviewerIDs {
		wg.Go(func() {
			_, err := tx.Exec(
				"INSERT INTO review_reviewers (review_id, reviewer_id) VALUES ($1, $2)",
				reviewID, reviewerID,
			)
			errChan <- err
		})
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(errChan)

	// Check for errors from the goroutines
	for err := range errChan {
		if err != nil {
			_ = tx.Rollback()
			http.Error(w, "Error adding reviewers", http.StatusInternalServerError)
			return
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		http.Error(w, "Error committing transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetReviews godoc
// @Summary Get all reviews
// @Description Fetches all reviews along with reviewers
// @Tags Admin
// @Produce json
// @Success 200 {array} types.ReviewResponse
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/reviews [get]
func GetReviews(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Conn.Query(`
		SELECT r.id, r.employee_id, e.email AS employee_email, r.performance_review, r.comments, r.created_at,
		       ARRAY_AGG(rr.reviewer_id) AS reviewer_ids
		FROM reviews r
		JOIN employees e ON r.employee_id = e.id
		LEFT JOIN review_reviewers rr ON r.id = rr.review_id
		GROUP BY r.id, e.email
	`)
	if err != nil {
		http.Error(w, "Error fetching reviews", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

	var reviews []types.ReviewResponse
	for rows.Next() {
		var id, employeeID int
		var employeeEmail, performanceReview string
		var comments []string
		var reviewerIDs []int
		var createdAt string
		err := rows.Scan(&id, &employeeID, &employeeEmail, &performanceReview, pq.Array(&comments), &createdAt, pq.Array(&reviewerIDs))
		if err != nil {
			http.Error(w, "Error scanning review data", http.StatusInternalServerError)
			return
		}
		reviews = append(reviews, types.ReviewResponse{
			ID:                id,
			EmployeeID:        employeeID,
			EmployeeEmail:     employeeEmail,
			PerformanceReview: performanceReview,
			Comments:          comments,
			ReviewerIDs:       reviewerIDs,
			CreatedAt:         createdAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reviews); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}
