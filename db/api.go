package db

func SetupConnection(redisAddress string) error {
	return setupConnection(redisAddress)
}

func AddEntity(key string, values map[string]string) error {
	return add(key, values)
}

func GetAllEntities(key string) ([]map[string]string, error) {
	return getAll(key)
}

func DeleteEntity(key, id string) (bool, error) {
	return deleteEntity(key, id)
}

func GetEntityById(key, id string) (string, error) {
	return getEntityById(key, id)
}

func UpdateEntity(key, id, value string) (bool, error) {
	return updateEntity(key, id, value)
}
