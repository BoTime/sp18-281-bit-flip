package main

import (
	"log"
	"context"
	"net/http"
)

var userIdKey = "user_id"

// Authenticate the User, add 'user_id' to request context
func (srv *Server) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authentication")

		user_id, found := getUser(token)
		if !found && srv.debug {
			// Allow passing of User ID directly if in debug mode
			user_id = r.Header.Get("X-User-ID")
		}

		if user_id != "" {
			// If User ID was found, continue as Authenticated User
			log.Printf("Authenticated User [%s]", user_id)
			user_context := context.WithValue(r.Context(), userIdKey, user_id)
			next.ServeHTTP(w, r.WithContext(user_context))
		} else {
			// Else throw error for Unauthenticated User
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

// Retrieve User ID from Authentication Token
func getUser(token string) (string, bool) {
	// TODO(bbamsch): Implement Token->UserID Logic
	return "", false
}

func GetUserId(r *http.Request) string {
	return r.Context().Value(userIdKey).(string)
}