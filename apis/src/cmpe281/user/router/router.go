/**
 * Author: Bo
 * Created at: April 16, 2018
 *
 * Routes and route handlers
 */

package router

import (
    "github.com/gorilla/mux"
    "github.com/urfave/negroni"
    "cmpe281/common"
    "cmpe281/user/handlers"
)

func InitRoutes() *negroni.Negroni {


    // ===== Default Router ===== //
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.HealthCheckHandler)
    router.HandleFunc("/signin", handlers.SigninPostHandler).Methods("POST")
    router.HandleFunc("/signin", handlers.SigninGetHandler).Methods("GET")
    router.HandleFunc("/signup", handlers.SignupHandler)
    router.HandleFunc("/logout", handlers.LogoutHandler)

    n := negroni.Classic()
    n.UseHandler(router)


    // ===== User Router ===== //
    userRouter := mux.NewRouter().PathPrefix("/users").Subrouter().StrictSlash(true)

    userRouter.HandleFunc("/", handlers.ListUsersHandler)
    userRouter.HandleFunc("/{userId}", handlers.GetUserHandler).Methods("GET")
    userRouter.HandleFunc("/{userId}", handlers.UpdateUserHandler).Methods("PUT")

    router.PathPrefix("/users").Handler(negroni.New(
      negroni.HandlerFunc(common.AuthMiddlewareNegroni),
      negroni.Wrap(userRouter),
    ))

    return n
}
