package database

import (
	"os"
	"rockyprabowo/assignment-2/helpers/options"
)

func Config() DbConfigMap {
	return DbConfigMap{
		"host":         options.Default(os.Getenv("DB_HOST"), "localhost"),
		"user":         options.Default(os.Getenv("DB_USER"), "postgres"),
		"password":     options.Default(os.Getenv("DB_PASSWORD"), "postgres"),
		"databaseName": options.Default(os.Getenv("DB_DATABASE"), "orders_by"),
		"sslMode":      options.Default(os.Getenv("DB_SSLMODE"), "disable"),
	}
}
