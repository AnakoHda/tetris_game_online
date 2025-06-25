package config

import (
	"fmt"
	"os"
)

func ParseEnvironment() error {
	massEnv := []string{
		"MAIL_USERNAME",
		"MAIL_PASSWORD",
		"MAIL_HOST",
		"MAIL_PORT",
		"KAFKA_BROKERS",
		"KAFKA_HALLO_TOPIC",
		"KAFKA_SCORE_UPDATE_TOPIC",
		"KAFKA_GROUP_ID",
	}
	for _, str := range massEnv {
		if os.Getenv(str) == "" {
			return fmt.Errorf("%s is not set", str)
		}
	}
	return nil
}
