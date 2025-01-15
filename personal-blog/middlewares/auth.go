package middleware

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the user is trying to log in
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		// Check for the session cookie
		cookie, err := r.Cookie("session")
		if err != nil || cookie.Value != "valid" {
			// If there's no valid session, redirect to login
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// If the session is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
