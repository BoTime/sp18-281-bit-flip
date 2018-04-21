package handlers

import (
    "net/http"
    _ "encoding/json"
    "fmt"
    _ "cmpe281/user/models"
)

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: BO
    // 1. Return a list of users
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "List all users success.")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    // TODO
    // 1. Get user by id
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Get user success.")
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: BO
    // 1. Write User info to database
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Update user success.")
}
