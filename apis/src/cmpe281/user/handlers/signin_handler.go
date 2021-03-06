package handlers

import (
    "net/http"
    "encoding/json"
    "fmt"
    "cmpe281/user/models"
    "cmpe281/common"
)

func SigninPostHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Put user authentication in a goroutine so that it won't block
    // allow cross domain AJAX requests
    // set response type to json
    // w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-type", "application/json")

    var (
        user models.User
        responseBody = make(map[string] interface{})
    )
    fmt.Println("request body====", r.Body)
    if err := json.NewDecoder(r.Body).Decode(&user); err == nil {
        // 1. Get userId
        // ignore error during dev phase
        if userId, err := user.VerifyPasswordAndReturnUserId(); err == nil && userId != "" {

            // 2. Issue jwt token
            if tokenString, err := common.IssueTokenForUserIdV2(userId); err == nil {
                w.Header().Set("Authorization", "jwt " + tokenString)
                w.WriteHeader(http.StatusOK)

                responseBody["msg"] = "Login success"
                responseBody["userId"] = userId
                responseBody["name"] = user.Name

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
        fmt.Println("Invalid JSON format")
        responseBody["msg"] = "Invalid JSON format"
    }

    // Send response
    body, _ := json.Marshal(responseBody)
    fmt.Fprintf(w, string(body))
}

func SigninGetHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "")
}
