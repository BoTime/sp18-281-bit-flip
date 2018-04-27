# User Backend

Starbuck User backend is responsible for user authentication and registration. JWT token is used for User authentication. Once user if verified, a signed JWT token will be attached to HTTP reponse Headers, in the format of `Authorization: jwt token-content`.

| Signing Method | HMAC Secret | Expiration |
| -------------- | ----------- | -----------|
| HS256          | bit-flip    | 300 seconds|


## API Reference
- [Base Url](#Base-Url)
- [Login](#login)
- [Signup](#signup)
- [Logout](#logout)
- [Get User](#get-user)

### Base Url
`kong`: http://kong-lb-133222058.us-west-1.elb.amazonaws.com/users/v1/

### Login

#### POST /users/v1/login
##### Request
`Header`

| Header | Description |
|--------|-------------|
| N/A | N/A |

`Body`

```json
{
  "email": string,
  "password": string
}
```

##### Response
`Status: 200 Ok`

`Header`

| Header | Description |
|--------|-------------|
| Authorization | jwt token-content |

`Body`

```json
{
    "msg": "Login success",
    "user_id": string
}
```

| Property Name | Type | Description |
|---------------|------|-------------|
| `user_id` | string | Format UUID v4 |

---

### Signup

#### POST /users/v1/signup
##### Request Headers
`Header`

| Header | Description |
|--------|-------------|
| N/A | N/A |

`Body`

```json
{
  "email": string,
  "password": string,
  "firstname": string,
  "lastname": string
}
```

##### Response
`Status: 200 Ok`

`Header`

| Header | Description |
|--------|-------------|
| Authorization | jwt token-content |

`Body`

```json
{
    "msg": "Signup success",
    "user_id": string
}
```

| Property Name | Type | Description |
|---------------|------|-------------|
| `user_id` | string | Format UUID v4 |


---

### Logout

#### POST /users/v1/logout
##### Request Headers
`Header`

| Header | Description |
|--------|-------------|
| N/A | N/A |

`Body`

```json
{
  "email": string,
  "password": string
}
```

##### Response
`Status: 200 Ok`

`Header`

| Header | Description |
|--------|-------------|
| N/A | N/A |

`Body`

```json
{
    "msg": "Logout success"
}
```

---

### Get User

#### GET /users/v1/user/{user_id}
##### Request Headers
`Header`

| Header | Description |
|--------|-------------|
| Authorization | jwt token-content |

`Body`

```json
```

##### Response
`Status: 200 Ok`

`Header`

| Header | Description |
|--------|-------------|
| Authorization | jwt token-content |

`Body`

```json
{
    "url": "http://kong-lb-133222058.us-west-1.elb.amazonaws.com/users/v1/{user_id}",
    "email": string,
    "firstname": string,
    "lastname": string,
    "user_id": string
}
```

---

### Update User

#### PATCH /users/v1/user/{user_id}
##### Request Headers
`Header`

| Header | Description |
|--------|-------------|
| Authorization | jwt token-content |

`Body`

```json
{
    "password": string
}
```

##### Response
`Status: 200 Ok`

`Header`

| Header | Description |
|--------|-------------|
| Authorization | jwt token-content |

`Body`

```json
{
    "url": "http://kong-lb-133222058.us-west-1.elb.amazonaws.com/users/v1/{user_id}",
    "email": string,
    "firstname": string,
    "lastname": string,
    "user_id": string
}
```

---

#### Resources

---
