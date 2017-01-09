package db

import (
	"log"

	redis "gopkg.in/redis.v5"
	"fmt"
)

var (
	redisClient *redis.Client
)

func SetupConnection(redisAddress string) error {
	log.Printf("Creating redis client with connection to %s.", redisAddress)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}
	if pong == "PONG" {
		log.Printf("Connection to redis database was successful.")
	} else {
		return fmt.Errorf("Ping to database returned unexpected value: '%s'.", pong)
	}
}
