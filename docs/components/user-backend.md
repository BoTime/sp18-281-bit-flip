# User Backend

Starbuck User backend is responsible for user authentication and registration. JWT token is used for User authentication. Once user if verified, a signed JWT token will be attached to HTTP reponse Headers, in the format of `Authorization: jwt token-content`.

| Signing Method | HMAC Secret | Expiration |
| -------------- | ----------- | -----------|
| HS256          | bit-flip    | 300 seconds|

- [Login](#login)
- [Signup](#signup)
- [Logout](#logout)
- [Get User](#get-user)

## API Reference

### Login

#### POST /starbucks/v1/login
##### Request Headers

---

### Signup

#### POST /starbucks/v1/signup
##### Request Headers

---

### Logout

#### POST /starbucks/v1/logout
##### Request Headers

---

### Get User

#### POST /starbucks/v1/user/{user_id}
##### Request Headers

---
