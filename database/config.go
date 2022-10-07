package database

import (
	"os"
	. "rocky.my.id/git/h8-assignment-2/helpers/options"
)

// DbConfigMap is an alias of a map[string]string.
type DbConfigMap map[string]any

// Config returns the database configurations.
func Config() DbConfigMap {
	return DbConfigMap{
		"host": Default(
			os.Getenv("DB_HOST"),
			"localhost",
		),
		"user": Default(
			os.Getenv("DB_USER"),
			"postgres",
		),
		"password": Default(
			os.Getenv("DB_PASSWORD"),
			"postgres",
		),
		"databaseName": Default(
			os.Getenv("DB_DATABASE"),
			"orders_by",
		),
		"sslMode": Default(
			os.Getenv("DB_SSLMODE"),
			"disable",
		),
	}
}
