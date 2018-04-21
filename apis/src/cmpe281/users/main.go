package main

import (
    "log"
    "net/http"
    "cmpe281/users/router"
    "os"
)

func main() {
    // Init routers
    n := router.InitRoutes()

    // default port server listens on
    port := "8080"
    if (os.Getenv("GO_SERVER_PORT") != "") {
        port = os.Getenv("GO_SERVER_PORT")
    }

    log.Printf("[*] Server started at port " + port)
    log.Fatal(http.ListenAndServe(":" + port, n))
}
