package main

import (
	"errors"
	"os"
)

func requiredEnvVars() []string {
	return []string{
		"APP_HOST",
		"APP_PORT",
		"APP_STATIC_URL",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_DATABASE",
		"DB_SSLMODE",
	}
}

func CheckEnvVars() error {
	for _, v := range requiredEnvVars() {
		if _, present := os.LookupEnv(v); !present {
			return errors.New("Environment variable " + v + " is not set!")
		}
	}
	return nil
}
