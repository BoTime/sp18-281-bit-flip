package handlers

import (
    "net/http"
    "fmt"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: BO
    // How does the frontend handle logout?
    w.WriteHeader(http.StatusOK)
    http.Redirect(w, r, "/signin", 300)
    fmt.Fprintf(w, "Logout success.")
}
