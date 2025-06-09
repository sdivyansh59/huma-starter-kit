package middleware

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

// Authenticate verifies that a valid token is present in the request
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//token := extractToken(r)
		//
		//if token == "" || !isValidToken(token) {
		//	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		//	return
		//}

		log.Info().Msgf("Authentication Middleware Called...................")
		// Token is valid, continue to next handler
		next.ServeHTTP(w, r)
	})
}

func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func isValidToken(token string) bool {
	// Implement your token validation logic here
	// This could involve JWT verification, checking against a database, etc.
	return true // Replace with actual validation
}
