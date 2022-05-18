package config

import (
	"log"
	"os"
)

func getEnvVar(key string) string {
	val, exist := os.LookupEnv(key)
	if !exist {
		log.Fatalf("no env var | %v", key)
	}
	return val
}
