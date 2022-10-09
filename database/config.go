package database

import (
	"errors"
	"net"
	"net/url"
	"os"
	. "rocky.my.id/git/h8-assignment-2/helpers/options"
	"strings"
)

// DbConfigMap is an alias of a map[string]string.
type DbConfigMap map[string]string

// Config returns the database configurations.
func Config() (config DbConfigMap, err error) {
	databaseURL, hasDatabaseURL := os.LookupEnv("DB_URL")
	if hasDatabaseURL {
		config, err = parseDatabaseURL(databaseURL)
		return
	}
	config = parseFromEnvVars()
	return
}

// parseFromEnvVars returns a database configuration from loaded environment variables.
func parseFromEnvVars() DbConfigMap {
	return DbConfigMap{
		"host": Default(
			os.Getenv("DB_HOST"),
			"localhost",
		),
		"port": Default(
			os.Getenv("DB_PORT"),
			"5432",
		),
		"user": Default(
			os.Getenv("DB_USER"),
			"postgres",
		),
		"password": Default(
			os.Getenv("DB_PASSWORD"),
			"postgres",
		),
		"dbname": Default(
			os.Getenv("DB_DATABASE"),
			"orders_by",
		),
		"sslmode": Default(
			os.Getenv("DB_SSLMODE"),
			"disable",
		),
	}
}

// parseDatabaseURL returns a database configuration from the database URL defined in DB_URL environment variable.
func parseDatabaseURL(databaseURL string) (config DbConfigMap, err error) {
	var parsedURL *url.URL
	var host, port, username string

	parsedURL, err = url.Parse(databaseURL)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	host, port, err = net.SplitHostPort(parsedURL.Host)
	username = parsedURL.User.Username()
	password, hasPassword := parsedURL.User.Password()

	config = DbConfigMap{
		"host":    host,
		"port":    port,
		"user":    username,
		"dbname":  strings.Trim(parsedURL.Path, "/"),
		"sslmode": Default(os.Getenv("DB_SSLMODE"), "disable"),
	}

	if hasPassword {
		config["password"] = password
	}

	return
}
