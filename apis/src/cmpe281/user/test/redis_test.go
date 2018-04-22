package test

import (
    "github.com/go-redis/redis"
    "github.com/joho/godotenv"
    "os"
    "testing"
)

func TestRedisClient(t *testing.T) {
    // Load env variables
    err := godotenv.Load("../.env")

    if err != nil {
        t.Error(err)
        return
    }

    REDIS_DOMAIN := os.Getenv("REDIS_DOMAIN")
    REDIS_PORT := os.Getenv("REDIS_PORT")

    t.Log("REDIS_DOMAIN", REDIS_DOMAIN)
    t.Log("REDIS_PORT", REDIS_PORT)
	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_DOMAIN + ":" + REDIS_PORT,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err = client.Ping().Result(); err != nil {
        t.Error(err)
    }
}

func TestRedisWrite(t *testing.T) {
    // Load env variables
    err := godotenv.Load("../.env")

    if err != nil {
        t.Error(err)
        return
    }

    REDIS_DOMAIN := os.Getenv("REDIS_DOMAIN")
    REDIS_PORT := os.Getenv("REDIS_PORT")

    t.Log("REDIS_DOMAIN", REDIS_DOMAIN)
    t.Log("REDIS_PORT", REDIS_PORT)
	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_DOMAIN + ":" + REDIS_PORT,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

    if err := client.Set("foo@bar.com", "bar", 0).Err(); err != nil {
        t.Error(err)
    }

    if val, err := client.Get("foo@bar.com").Result(); err == nil {
        if val != "bar" {
            t.Error("wrong value")
        }
    } else  {
        t.Error(err)
    }
}

func TestRedisRead(t *testing.T) {
    // Load env variables
    err := godotenv.Load("../.env")

    if err != nil {
        t.Error(err)
        return
    }

    REDIS_DOMAIN := os.Getenv("REDIS_DOMAIN")
    REDIS_PORT := os.Getenv("REDIS_PORT")

    t.Log("REDIS_DOMAIN", REDIS_DOMAIN)
    t.Log("REDIS_PORT", REDIS_PORT)
	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_DOMAIN + ":" + REDIS_PORT,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

    if val, err := client.Get("foo@bar.com").Result(); err == nil {
        if val != "bar" {
            t.Error("wrong value ")
        }
    } else  {
        t.Error(err)
    }
}
