package db

import (
	"fmt"
	"log"

	"gopkg.in/redis.v5"
	"strconv"
)

var (
	redisClient *redis.Client
)

func setupConnection(redisAddress string) error {
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
	return err
}

func add(key string, values map[string]string) (string, error) {
	idCmd := redisClient.Incr(fmt.Sprintf("%ss:lastid", key))
	if idCmd.Err() != nil {
		return "", idCmd.Err()
	}
	lastId, err := idCmd.Result()
	if err != nil {
		return "", err
	}
	lastIdStr := strconv.Itoa(int(lastId))
	newKey := fmt.Sprintf("%s:%s", key, lastIdStr)
	values["Id"] = lastIdStr
	result := redisClient.HMSet(newKey, values)
	if result.Err() != nil {
		return "", result.Err()
	}
	_, err = result.Result()
	if err != nil {
		return "", err
	}
	saddCmd := redisClient.SAdd(fmt.Sprintf("%ss", key), newKey)
	if saddCmd.Err() != nil {
		return lastIdStr, saddCmd.Err()
	}
	_, err = saddCmd.Result()
	if err != nil {
		return lastIdStr, err
	}
	return lastIdStr, nil
}

func getAll(key string) ([]map[string]string, error) {
	keys, err := getKeysSet(key)
	if err != nil {
		return nil, err
	}
	entities := []map[string]string{}
	for i := 0; i < len(keys); i++ {
		allCmd := redisClient.HGetAll(keys[i])
		if allCmd.Err() != nil {
			return nil, allCmd.Err()
		}
		result, err := allCmd.Result()
		if err != nil {
			return nil, err
		}
		entities = append(entities, result)
	}
	return entities, nil
}

func deleteEntity(key, id string) (bool, error) {
	entityKey := fmt.Sprintf("%s:%s", key, id)
	remCmd := redisClient.SRem(fmt.Sprintf("%ss", key), entityKey)
	if remCmd.Err() != nil {
		return false, remCmd.Err()
	}
	result, err := remCmd.Result()
	if err != nil {
		return false, err
	}
	if int(result) != 1 {
		return false, nil
	}
	delCmd := redisClient.Del(entityKey)
	if delCmd.Err() != nil {
		return false, delCmd.Err()
	}
	result, err = delCmd.Result()
	if err != nil {
		return false, err
	}
	return int(result) == 1, nil
}

func getEntityById(key, id string) (map[string]string, error) {
	entityKey := fmt.Sprintf("%s:%s", key, id)
	exists, err := keyExists(entityKey)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, nil
	}
	getCmd := redisClient.HGetAll(entityKey)
	if getCmd.Err() != nil {
		return nil, getCmd.Err()
	}
	result, err := getCmd.Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func updateEntity(key, id string, values map[string]string) error {
	entityKey := fmt.Sprintf("%s:%s", key, id)
	result := redisClient.HMSet(entityKey, values)
	if result.Err() != nil {
		return result.Err()
	}
	_, err := result.Result()
	if err != nil {
		return err
	}
	return nil
}

func entityExists(key, id string) (bool, error) {
	entityKey := fmt.Sprintf("%s:%s", key, id)
	return keyExists(entityKey)
}

func keyExists(key string) (bool, error) {
	exCmd := redisClient.Exists(key)
	if exCmd.Err() != nil {
		return false, fmt.Errorf("Verification of %ss value failed. %v", key, exCmd.Err())
	}
	exists, err := exCmd.Result()
	if err != nil {
		return false, fmt.Errorf("Processing of %ss verification failed. %v", key, err)
	}
	return exists, nil
}

func getKeysSet(key string) ([]string, error) {
	plKey := fmt.Sprintf("%ss", key)
	exists, err := keyExists(plKey)
	if err != nil {
		return nil, err
	}
	if !exists {
		return []string{}, nil
	}
	smemCmd := redisClient.SMembers(fmt.Sprintf("%ss", key))
	if smemCmd.Err() != nil {
		return nil, fmt.Errorf("Retrieval of %ss set failed. %v", key, smemCmd.Err())
	}
	keys, err := smemCmd.Result()
	if err != nil {
		return nil, fmt.Errorf("Processing of %ss set retrieval failed. %v", key, err)
	}
	return keys, nil
}
