/**
 * Author: Bo
 */

// Generate and validate JWT tokens
package common

import (
    "fmt"
    jwt "github.com/dgrijalva/jwt-go"
    "github.com/satori/go.uuid"
    _ "time"
    "encoding/json"
)

var hmacSecret = "bit-flip"

// Parse JWT token using jwt.SigningMethodHMAC
func ParseToken(tokenString string) (*jwt.Token, error) {
    // Don't forget to validate the alg is what you expect:
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Verify signing algorithm
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }

		return []byte(hmacSecret), nil
	})


    return token, err
}

// Generate JWT token from given json string
func IssueToken(jsonPayload []byte) (string, error) {
    // Parse json to jwt.MapClaims (of type map[string] interface{})
    var mapClaims jwt.MapClaims
    if err := json.Unmarshal(jsonPayload, &mapClaims); err != nil {
        return "", err
    }

    // Create a new token object, specifying signing method and the claims
    // you would like it to contain.
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

    // Sign and get the complete encoded token as a string using the secret
    tokenString, err := token.SignedString([]byte(hmacSecret))

    return tokenString, err
}


func IssueTokenForUserId(userId *uuid.UUID) (string, error) {
    jsonString := `{"user_id":"` + userId.String() + `"}`
    tokenString, err := IssueToken([]byte(jsonString))
    return tokenString, err
}

func IssueTokenForUserIdV2(userId string) (string, error) {
    jsonString := `{"user_id":"` + userId + `"}`
    tokenString, err := IssueToken([]byte(jsonString))
    return tokenString, err
}
