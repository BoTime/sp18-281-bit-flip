package handlers

import (
    "net/http"
    "encoding/json"
    "fmt"
    "cmpe281/users/models"
    "cmpe281/common"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Put user authentication in a goroutine so that it won't block
    // allow cross domain AJAX requests
    // set response type to json
    // w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-type", "application/json")

    var (
        user models.User
        responseBody = make(map[string] interface{})
    )

    if err := json.NewDecoder(r.Body).Decode(&user); err == nil {
        // 1. Get userId
        // ignore error during dev phase
        if userId, err := user.VerifyPasswordAndReturnUserId(); err == nil && userId != nil {

            // 2. Issue jwt token
            if tokenString, err := common.IssueTokenForUserId(userId); err == nil {
                w.Header().Set("Authorization", "jwt " + tokenString)
                w.WriteHeader(http.StatusOK)

                responseBody["msg"] = "Login success"
                responseBody["userId"] = userId

            } else {
                // [Error] Unable to geenrate JWT token
                w.WriteHeader(http.StatusInternalServerError)
                responseBody["msg"] = err.Error()
            }

        } else {
            // [Error] Invalid Email or Password
            w.WriteHeader(http.StatusUnauthorized)
            responseBody["msg"] = err.Error()
        }

    } else {
        // [Error] Bad json format
        w.WriteHeader(http.StatusBadRequest)
        responseBody["msg"] = "Invalid JSON format"
    }

    // Send response
    body, _ := json.Marshal(responseBody)
    fmt.Fprintf(w, string(body))
}
