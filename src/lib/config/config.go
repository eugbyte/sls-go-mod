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
	AWS_DEFAULT_REGION    string
	AWS_ACCESS_KEY_ID     string
	AWS_SECRET_ACCESS_KEY string
	DYNAMO_DB_ENDPOINT    string
}

var Config = conf{
	AWS_DEFAULT_REGION:    getEnvOrDefault("AWS_DEFAULT_REGION", "ap-southeast-1"),
	AWS_ACCESS_KEY_ID:     getEnvOrDefault("AWS_ACCESS_KEY_ID", "123"),
	AWS_SECRET_ACCESS_KEY: getEnvOrDefault("AWS_SECRET_ACCESS_KEY", "123"),
	DYNAMO_DB_ENDPOINT:    getEnvOrDefault("DYNAMO_DB_ENDPOINT", "http://host.docker.internal:18000"),
}
