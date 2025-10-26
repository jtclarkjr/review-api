package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"go-api/db"
	"go-api/types"
)

type contextKey string

const userIDKey contextKey = "id"

// ListReviews godoc
// @Summary List assigned reviews
// @Description Lists reviews assigned to the employee that have not been submitted yet
// @Tags Employee
// @Produce json
// @Success 200 {array} types.AssignedReviewResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /employee/reviews [get]
func ListReviews(w http.ResponseWriter, r *http.Request) {
	// Get id from context added from auth
	idValue := r.Context().Value(userIDKey)
	employeeID, ok := idValue.(string)
	if !ok || employeeID == "" {
		http.Error(w, "Invalid ID", http.StatusUnauthorized)
		return
	}

	// Fetch reviews assigned to the employee that have not been submitted yet
	rows, err := db.Conn.QueryContext(r.Context(), `
        SELECT r.id, e.email AS employee_email, r.performance_review
        FROM reviews r
        JOIN employees e ON r.employee_id = e.id
        LEFT JOIN feedback f ON r.id = f.review_id AND f.submitted = TRUE
        WHERE f.review_id IS NULL AND r.id IN (
            SELECT review_id FROM review_reviewers WHERE reviewer_id = $1
        )
    `, employeeID)
	if err != nil {
		http.Error(w, "Error fetching reviews", http.StatusInternalServerError)
		return
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			log.Printf("Error closing rows: %v", closeErr)
		}
	}()

	// Build the list of reviews
	var reviews []types.AssignedReviewResponse
	for rows.Next() {
		var id int
		var employeeEmail, performanceReview string
		if err := rows.Scan(&id, &employeeEmail, &performanceReview); err != nil {
			http.Error(w, "Error scanning review data", http.StatusInternalServerError)
			return
		}
		reviews = append(reviews, types.AssignedReviewResponse{
			ID:                id,
			EmployeeEmail:     employeeEmail,
			PerformanceReview: performanceReview,
		})
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating over review data", http.StatusInternalServerError)
		return
	}

	// Respond with the list of reviews
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reviews); err != nil {
		http.Error(w, "Error encoding response for listed reviews", http.StatusInternalServerError)
	}
}

// SubmitFeedback godoc
// @Summary Submit review feedback
// @Description Submits feedback for a review assigned to the employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param feedback body object true "Feedback info"
// @Success 201 {object} types.MessageResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /employee/reviews/feedback [post]
func SubmitFeedback(w http.ResponseWriter, r *http.Request) {
	// Get id from context added from auth
	idValue := r.Context().Value(userIDKey)
	employeeID, ok := idValue.(string)
	if !ok || employeeID == "" {
		http.Error(w, "Invalid ID", http.StatusUnauthorized)
		return
	}

	// Parse the incoming JSON payload
	var feedback struct {
		ReviewID int    `json:"review_id"`
		Comment  string `json:"comment"`
	}

	defer func() {
		if closeErr := r.Body.Close(); closeErr != nil {
			log.Printf("Error closing request body: %v", closeErr)
		}
	}()
	if err := json.NewDecoder(r.Body).Decode(&feedback); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate that the employee is authorized to review the given review
	var isReviewer bool
	err := db.Conn.QueryRowContext(r.Context(),
		"SELECT EXISTS(SELECT 1 FROM review_reviewers WHERE review_id = $1 AND reviewer_id = $2)",
		feedback.ReviewID, employeeID,
	).Scan(&isReviewer)
	if err != nil {
		http.Error(w, "Error validating reviewer status", http.StatusInternalServerError)
		return
	}

	if !isReviewer {
		http.Error(w, "Unauthorized to review this performance review", http.StatusForbidden)
		return
	}

	// Append the feedback to the comments in the reviews table
	_, err = db.Conn.ExecContext(r.Context(),
		"UPDATE reviews SET comments = array_append(comments, $1) WHERE id = $2",
		feedback.Comment, feedback.ReviewID,
	)
	if err != nil {
		http.Error(w, "Error adding feedback to review", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(types.MessageResponse{
		Message: "Feedback submitted successfully",
	}); err != nil {
		http.Error(w, "Error submitting feedback", http.StatusInternalServerError)
	}
}
