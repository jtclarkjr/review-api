package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go-api/db"
	"go-api/types"
	"golang.org/x/crypto/bcrypt"
)

// JwtKey is the secret key used to sign tokens
var JwtKey = []byte("your_secret_key")

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

// Login godoc
// @Summary Login to generate a JWT token
// @Description Logs in a user with email and password, and returns a JWT token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body handlers.Credentials true "Email and Password"
// @Success 200 {object} types.TokenResponse
// @Failure 401 {string} string "Unauthorized"
// @Router /login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var userID int
	var hashedPassword, role string
	err = db.Conn.QueryRow("SELECT id, password, role FROM users WHERE email = $1", creds.Email).Scan(&userID, &hashedPassword, &role)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID:    userID,
		Email: creds.Email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := types.TokenResponse{Token: tokenString}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("Error encoding response JSON: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// hashPassword generates a bcrypt hash for the given password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// ExtractClaims extracts the claims from the Authorization header in the HTTP request
func ExtractClaims(r *http.Request) (*Claims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header missing")
	}

	// Expecting "Bearer <token>" format
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("invalid authorization header format")
	}

	// Extract the token string
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse the token and validate it
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}
