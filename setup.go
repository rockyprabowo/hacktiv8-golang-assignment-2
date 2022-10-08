package main

import (
	"errors"
	"log"
	"os"
	"strings"
)

func requiredEnvVars() []string {
	var vars = []string{
		"APP_HOST",
		"APP_PORT",
		"APP_STATIC_URL",
		"DB_SSLMODE",
	}

	if _, present := os.LookupEnv("DB_URL"); present {
		log.Println("config: using DB_URL configurations")
		vars = append(vars, "DB_URL")
	} else {
		log.Println("config: using DB_* configurations")
		vars = append(vars, []string{
			"DB_HOST",
			"DB_PORT",
			"DB_USER",
			"DB_PASSWORD",
			"DB_DATABASE",
		}...)
	}

	return vars
}

func CheckEnvVars() error {
	var undefinedEnvVars []string
	for _, v := range requiredEnvVars() {
		if _, present := os.LookupEnv(v); !present {
			undefinedEnvVars = append(undefinedEnvVars, v)
		}
	}

	if len(undefinedEnvVars) > 0 {
		return errors.New(
			"Environment variable " +
				strings.Join(undefinedEnvVars, ", ") +
				" is not set!",
		)
	}
	return nil
}
