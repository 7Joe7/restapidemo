package db

import (
	"fmt"
	"log"

	redis "gopkg.in/redis.v5"
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
		return nil
	}
	return fmt.Errorf("Ping to database returned unexpected value: '%s'.", pong)
}

func Add(key, p string) error {
	idCmd := redisClient.Incr(fmt.Sprintf("%s:lastid", key))
	if idCmd.Err() != nil {
		return idCmd.Err()
	}
	lastId, err := idCmd.Result()
	if err != nil {
		return err
	}
	result := redisClient.HSet(key, string(lastId), p)
	if result.Err() != nil {
		return result.Err()
	}
	r, err := result.Result()
	if err != nil {
		return fmt.Errorf("Add failed with error: %v and message: %s.", err, result.String())
	}
	if !r {
		return fmt.Errorf("Add failed. %s", result.String())
	}
	return nil
}

func GetAll(key string) (map[string]string, error) {
	allCmd := redisClient.HGetAll(key)
	if allCmd.Err() != nil {
		return nil, allCmd.Err()
	}
	result, err := allCmd.Result()
	if err != nil {
		return nil, fmt.Errorf("Get All failed with error: %v and message: %s.", allCmd.Err(), allCmd.String())
	}
	return result, nil
}

func Delete(key, id string) {

}

func GetById(key, id string) (string, error) {
	getCmd := redisClient.HGet(key, id)
	if getCmd.Err() != nil {
		return nil, getCmd.Err()
	}
	result, err := getCmd.Result()
	if err != nil {
		return nil, fmt.Errorf("Get failed with error: %v and message: %s.", getCmd.Err(), getCmd.String())
	}
	return result, nil
}
