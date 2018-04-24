package models

import (
    "github.com/satori/go.uuid"
    "github.com/go-redis/redis"
    "github.com/joho/godotenv"
    "encoding/json"
    "time"
    "errors"
    "log"
    "os"
)

type User struct {
    Email string `json:"email" binding:"required"`
    Password string `json:"password"`
    UserId string `json:"user_id"`
}

var (
    emailMock = "foo@bar.com"
    passwordMock = "bar"
)


// ===== Connect to Redis DB ===== //
var client = connectRedis()

func connectRedis() *redis.Client {

    var (
        REDIS_DOMAIN string
        REDIS_PORT string
    )

    if err := godotenv.Load("../.env"); err != nil {
        REDIS_DOMAIN = "localhost"
        REDIS_PORT = "6379"

    } else {
        REDIS_DOMAIN = os.Getenv("REDIS_DOMAIN")
        REDIS_PORT = os.Getenv("REDIS_PORT")
    }

    client := redis.NewClient(&redis.Options{
        Addr:     REDIS_DOMAIN + ":" + REDIS_PORT,
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    log.Println("[*] Connecte to Redis server at ", REDIS_DOMAIN + ":" + REDIS_PORT)
    if _, err := client.Ping().Result(); err != nil {
        log.Println("[x] Unable to start Redis server")
        panic("Redis")
    }

    return client
}


// ===== Public APIs ===== //
func (u User) VerifyPasswordAndReturnUserId() (string, error) {
    if user, err := getUserByEmail(u.Email); err != nil {
        // Probably server error
        log.Println(err)
        return "", err

    } else {
        if user != nil && passwordMatches(user.Password, u.Password) == true {
            // Login success
            return user.UserId, nil

        } else {
            // Login failed
            return "", errors.New("Invalid Email or Password")
        }
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

func (u User) CreateUserId(newUser *User) (*uuid.UUID, error) {
    // TODO: (Bo)
    // 1. Check email fields and password fields
    // make sure they are not empty
    if newUser.Email == "" || newUser.Password == "" {
        return nil, errors.New("Email and password cannot be empty")
    }

    // 2. Make sure email is not used
    log.Println("try to register === ", newUser.Email, "password: ", newUser.Password)
    userFound, err := client.Get(newUser.Email).Result()
    if err != nil && err.Error() != "redis: nil" {
        log.Println("[* | Sign Up] Failed to read from Redis")
        log.Println(err)
        return nil, errors.New("Server is busy, please try again.")
    }

    var userId uuid.UUID

    if userFound != "" {
        // email used
        return nil, errors.New("Email address already used by others.")

    } else {
        // email not used
        userId = uuid.Must(uuid.NewV4())
        newUser.UserId = userId.String()
        userString, _ := json.Marshal(newUser)

        log.Println("useString === ", string(userString))

        if err := client.Set(newUser.Email, userString, 0).Err(); err != nil {
            log.Println("[* | Sign Up] Failed to write email to Redis")
            log.Println(err)
            return nil, errors.New("Server is busy, please try again.")
        }

        if err := client.Set(newUser.UserId, userString, 0).Err(); err != nil {
            log.Println("[* | Sign Up] Failed to write userId to Redis")
            log.Println(err)
            return nil, errors.New("Server is busy, please try again.")
        }
    }

    return &userId, nil
}


// ===== Helper Functions ===== //

func passwordMatches(hashedPassword string, rawPassword string) bool {
    if hashedPassword == rawPassword {
        return true
    } else {
        return false
    }
}

// TODO: Implement Hashing
func getHashedString(input string) string {
    return input
}

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

func getUserByEmail(email string) (*User, error) {
    userInfoString, err := client.Get(email).Result()
    if err != nil && err.Error() == "redis: nil" {
        // Email does not exists
        log.Println("[* | Sign Up] Failed to read from Redis")
        log.Println(err)
        return nil, errors.New("Invalid Email or Password.")

    } else {
        var user User
        log.Println("user string ==== ", userInfoString)
        if err := json.Unmarshal([]byte(userInfoString), &user); err == nil {
            return &user, nil
        } else {
            // JSON parse error
            return nil, err
        }
    }
}
