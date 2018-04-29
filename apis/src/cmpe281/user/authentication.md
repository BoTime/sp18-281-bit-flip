## User authentication

### Why do we want to do authentication?
Authentication almost makes any application slower, but without it an application probably won't work for too long.

### What tools do we use to do authentication?
JSON Web Token (JWT) is widely used for authentication and information exchange.

If you want to know more about JWT, this article ([link](https://auth0.com/docs/jwt)) from auth0 is somewhere to start with.

If you want to play around with JWT online or check whether your JWT middleware gives the correct result, go to [here](https://jwt.io/).

### What library is used?

In this project, we used an Golang implementation of JWT from the following open source library:

https://github.com/dgrijalva/jwt-go

### Ok, show me some code...
For anyone who want to write less code:
```Golang
import
{
    cmpe281/common
}
```

For `github.com/gorilla/mux` users, authentication middleware can be added to router as follows:
```Golang
debug = true
router := mux.NewRouter()
router.Use(common.AuthMiddleware(debug))
```

Refer to this link for implementation:<br>
https://github.com/nguyensjsu/team281-bit-flip/blob/master/apis/src/cmpe281/payment/server.go


For `https://github.com/urfave/negroni` users, it is a little bit different:
```Golang
userRouter := mux.NewRouter().PathPrefix("/users").Subrouter().StrictSlash(true)
router.PathPrefix("/users").Handler(negroni.New(
    negroni.HandlerFunc(common.AuthMiddlewareNegroni),
    negroni.Wrap(userRouter),
))
```
Refer to this link for implementation:<br>
https://github.com/nguyensjsu/team281-bit-flip/blob/master/apis/src/cmpe281/user/router/router.go

For anyone doesn't do authentication, good luck~
