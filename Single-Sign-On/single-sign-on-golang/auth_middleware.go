package main

import (
	"net/http"
	"strings"
)

// AuthMiddleware verifies the token in the request header
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		// Check if the token exists in our fake user database
		isValid := false
		for _, user := range users {
			if user.Token == token {
				isValid = true
				break
			}
		}

		if !isValid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// If token is valid, continue processing request
		next.ServeHTTP(w, r)
	})
}
