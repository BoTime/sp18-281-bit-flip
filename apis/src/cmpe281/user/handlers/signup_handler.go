package handlers

import (
    "net/http"
    "encoding/json"
    "fmt"
    "cmpe281/user/models"
    "cmpe281/common"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")

    var (
        user models.User
        responseBody = make(map[string] interface{})
    )
    fmt.Println("request body====", r.Body)
    fmt.Println("request header Auth ====", r.Header.Get("Authorization"))
    if err := json.NewDecoder(r.Body).Decode(&user); err == nil {
        // 1. Create a new user
        if userId, err := user.CreateUserId(&user); err == nil && userId != nil {

            // 2. Issue jwt token
            if tokenString, err := common.IssueTokenForUserIdV2(userId.String()); err == nil {
                w.Header().Set("Authorization", "jwt " + tokenString)
                w.WriteHeader(http.StatusOK)

                responseBody["msg"] = "Signup success"
                responseBody["userId"] = userId
                responseBody["name"] = user.Name

            } else {
                // [Error] Unable to geenrate JWT token
                w.WriteHeader(http.StatusInternalServerError)
                responseBody["msg"] = err.Error()
            }

        } else {
            // [Error] Email used by other
            w.WriteHeader(http.StatusBadRequest)
            responseBody["msg"] = err.Error()
        }

    } else {
        // [Error] Bad json format
        w.WriteHeader(http.StatusBadRequest)
        fmt.Println("Bad json format")
        responseBody["msg"] = "Invalid JSON format"
    }

    fmt.Println(responseBody["msg"])
    body, _ := json.Marshal(responseBody)
    fmt.Fprintf(w, string(body))
}
