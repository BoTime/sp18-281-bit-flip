/**
 * Author: Bo
 * Created at: April 16, 2018
 */

 package router

 import (
 	"log"
 	"net/http"
 )

 // NewServer configures and returns a Server.
func main() {
    // Init route handler
    router := InitRoutes()

    port := "3001"

    log.Printf("[*] Server started at port " + port)
    http.ListenAndServe(":" + port, router)
}
