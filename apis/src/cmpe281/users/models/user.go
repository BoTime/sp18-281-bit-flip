package models

import (
    "github.com/satori/go.uuid"
    "time"
    "errors"
    "fmt"
)

type User struct {
    Email string `json:"email" binding:"required"`
    Password string `json:"password"`
}

var (
    emailMock = "foo@bar.com"
    passwordMock = "bar"
)

func (u User) VerifyPasswordAndReturnUserId() (*uuid.UUID, error) {
    var userId *uuid.UUID
    fmt.Println("====", u.Email, u.Password)
    if passwordIsValid(u.Email, u.Password) == true {
        userId, _ = getUserIdByEmail(u.Email)
        return userId, nil

    } else {
        return nil, errors.New("Invalid Email or Password")
    }
}

func (u User) FindById(userId uuid.UUID) *User {
    return &User{}
}

func (u User) UpdateById(userId uuid.UUID) bool {
    return true
}

func (u User) DeleteById(userId uuid.UUID) bool {
    return true
}

func (u User) CreateUserId(user *User) (*uuid.UUID, error) {
    // TODO: (Bo)
    // 1. Check email fields and password fields
    // make sure they are not empty
    if user.Email == "" || user.Password == "" {
        return nil, errors.New("Email and password cannot be empty")
    }

    // 2. Make sure email is not used
    if user.Email != emailMock {
        id := uuid.Must(uuid.NewV4())
        return &id, nil
    }

    return nil, errors.New("Email address already used by others.")

    // 3. Generate userId
    // (optional) Make sure userId is not used too

    // 4. Write User information to databse


}


// ===== Helper Functions ===== //

func passwordIsValid(email string, password string) bool {
    // TODO: Connect to DB to verify email and password
    // Simulate db query by sleep for 3 seconds
    time.Sleep(time.Millisecond * 100)
    return email == emailMock && password == passwordMock
}

func getUserIdByEmail(email string) (*uuid.UUID, error) {
    // TODO: Connect to DB to retrieve userId
    // error is ignored here
    time.Sleep(time.Millisecond * 100)
    id := uuid.Must(uuid.NewV4())
    return &id, nil
}
