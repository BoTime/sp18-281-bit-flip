package main

import (
	"log"
	"context"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"strings"
)

var userIdKey = "user_id"

// Authenticate the User, add 'user_id' to request context
func (srv *Server) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

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

func getUser(rawToken string) (string, bool) {
	fields := strings.Fields(rawToken)
	if len(fields) < 1 {
		log.Println("User Auth Token Missing")
		return "", false
	}

	switch fields[0] {
	case "jwt":
		return getUserJwt(fields)
	default:
		log.Println("Unexpected User Auth Token Type")
		return "", false
	}
}

// Retrieve User ID from Authentication Token
func getUserJwt(fields []string) (string, bool) {
	if len(fields) != 2 {
		log.Println("Invalid JWT Token")
		return "", false
	}

	token, err := jwt.Parse(fields[1], func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("bit-flip"), nil
	})

	if err != nil {
		log.Println("User Auth Token failed Parsing")
		return "", false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["user_id"]
		if userId != nil {
			return userId.(string), true
		}
		log.Println("User ID not in Claims")
		return "", false
	} else {
		log.Println("User Auth Token failed Validation")
		return "", false
	}
}

func GetUserId(r *http.Request) string {
	return r.Context().Value(userIdKey).(string)
}
