package utils

import (
	"log"
	"os"
)

func GetEnv(key string) string {
	return TrimFull(os.Getenv(key))
}

func GetEnvMust(key string) string {
	val := GetEnv(key)
	if val == EMPTY_STRING {
		log.Fatalf("missing environment variable: %v", key)
	}

	return val
}
