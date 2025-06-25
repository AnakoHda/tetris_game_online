package config

import (
	"fmt"
	"os"
)

func ParseEnvironment() error {
	massEnv := []string{
		"DB_USER",
		"DB_PASSWORD",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_HOST",
		"DB_PORT",
		"DB_URL",
		"JWT_SECRET",
		"KAFKA_BROKERS",
		"KAFKA_HALLO_TOPIC",
	}
	for _, str := range massEnv {
		if os.Getenv(str) == "" {
			return fmt.Errorf("%s is not set", str)
		}
	}
	return nil
}
