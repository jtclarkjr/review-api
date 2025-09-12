package types

// EmployeeResponse represents an employee in API responses
type EmployeeResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Position string `json:"position"`
}

// ReviewResponse represents a review in API responses
type ReviewResponse struct {
	ID                int      `json:"id"`
	EmployeeID        int      `json:"employee_id"`
	EmployeeEmail     string   `json:"employee_email"`
	PerformanceReview string   `json:"performance_review"`
	Comments          []string `json:"comments"`
	ReviewerIDs       []int    `json:"reviewer_ids"`
	CreatedAt         string   `json:"created_at"`
}

// AssignedReviewResponse represents a review assigned to an employee
type AssignedReviewResponse struct {
	ID                int    `json:"id"`
	EmployeeEmail     string `json:"employee_email"`
	PerformanceReview string `json:"performance_review"`
}

// CreateEmployeeResponse represents the response when creating an employee
type CreateEmployeeResponse struct {
	EmployeeID int    `json:"employee_id"`
	Email      string `json:"email"`
}

// CreateReviewResponse represents the response when creating a review
type CreateReviewResponse struct {
	ReviewID int `json:"review_id"`
}

// MessageResponse represents a simple message response
type MessageResponse struct {
	Message string `json:"message"`
}

// TokenResponse represents a JWT token response
type TokenResponse struct {
	Token string `json:"token"`
}