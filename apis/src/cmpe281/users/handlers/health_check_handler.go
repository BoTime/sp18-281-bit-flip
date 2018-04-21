package handlers

import (
    "net/http"
    "fmt"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Starbucks User API V1.0 - Bo")
}
