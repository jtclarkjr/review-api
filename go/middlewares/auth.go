package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"go-api/handlers"
)

type contextKey string

const userIDKey contextKey = "id"

// AuthAdmin is middlewares that validates a JWT token and ensures the user has an admin role
func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if !strings.HasPrefix(tokenStr, "Bearer ") {
			http.Error(w, "Unauthorized: missing or invalid token", http.StatusUnauthorized)
			return
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		claims := &handlers.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return handlers.JwtKey, nil // Use the proper JwtKey from handlers
	})

		if err != nil {
			fmt.Printf("Error while parsing token: %v\n", err)
		}

		if !token.Valid || claims.Role != "admin" {
			http.Error(w, "Forbidden: invalid token or insufficient privileges", http.StatusForbidden)
			return
		}

		// Pass the email to the request context
		// Maybe context not needed anymore
		ctx := context.WithValue(r.Context(), userIDKey, claims.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AuthEmployee is middlewares that validates a JWT token and ensures the user has an employee role
func AuthEmployee(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if !strings.HasPrefix(tokenStr, "Bearer ") {
			http.Error(w, "Unauthorized: missing or invalid token", http.StatusUnauthorized)
			return
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
		claims := &handlers.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return handlers.JwtKey, nil
	})
		if err != nil || !token.Valid || claims.Role != "employee" {
			http.Error(w, "Forbidden: invalid token or insufficient privileges", http.StatusForbidden)
			return
		}

		// Pass the email to the request context
		ctx := context.WithValue(r.Context(), userIDKey, claims.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
