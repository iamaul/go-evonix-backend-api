package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)

var (
	Store *redis.Client
	err   error
)

func ConnectRedis() {
	Store = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := Store.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[redis] Connected successfully.")
}
