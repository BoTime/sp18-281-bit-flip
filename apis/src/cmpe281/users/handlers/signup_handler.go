package handlers

import (
    "net/http"
    "encoding/json"
    "fmt"
    "cmpe281/users/models"
    "cmpe281/common"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")

    var (
        user models.User
        responseBody = make(map[string] interface{})
    )

    fmt.Println("==========")
    if err := json.NewDecoder(r.Body).Decode(&user); err == nil {
        fmt.Println("+++++++")
        // 1. Create a new user
        if userId, err := user.CreateUserId(&user); err == nil && userId != nil {

            // 2. Issue jwt token
            if tokenString, err := common.IssueTokenForUserId(userId); err == nil {
                w.Header().Set("Authorization", "jwt " + tokenString)
                w.WriteHeader(http.StatusOK)

                responseBody["msg"] = "Signup success"
                responseBody["userId"] = userId

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
        responseBody["msg"] = "Invalid JSON format"
    }


    body, _ := json.Marshal(responseBody)
    fmt.Fprintf(w, string(body))
}
