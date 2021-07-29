package config

import "os"

func getEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

type conf struct {
	REGION             string
	ACCESS_KEY_ID      string
	SECRET_ACCESS_KEY  string
	DYNAMO_DB_ENDPOINT string
}

var Config = conf{
	REGION:             getEnvOrDefault("REGION", "ap-southeast-1"),
	ACCESS_KEY_ID:      getEnvOrDefault("ACCESS_KEY_ID", "123"),
	SECRET_ACCESS_KEY:  getEnvOrDefault("SECRET_ACCESS_KEY", "123"),
	DYNAMO_DB_ENDPOINT: getEnvOrDefault("DYNAMO_DB_ENDPOINT", "http://host.docker.internal:18000"),
}
