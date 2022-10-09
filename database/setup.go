package database

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"rocky.my.id/git/h8-assignment-2/models"
	"strings"
)

// buildDSN builds the DSN string from a database configuration.
func buildDSN(config DbConfigMap) (string, error) {
	var sb strings.Builder

	fields := []string{"host", "user", "dbname", "sslmode"}

	if _, hasPassword := config["password"]; hasPassword {
		fields = append(fields, "password")
	}

	for _, v := range fields {
		value, exist := config[v]
		{
			if !exist {
				return "", errors.New("Mandatory database configuration " + v + " is missing!")
			}
			sb.WriteString(v + "=" + value + " ")
		}
	}

	return strings.Trim(sb.String(), " "), nil
}

// DSN returns the DSN of the database connection.
func DSN() (string, error) {
	config, err := Config()
	if err != nil {
		return "", err
	}
	return buildDSN(config)
}

// Init initialise the database connection.
func Init() (db *gorm.DB) {
	var (
		dsn string
		err error
	)

	dsn, err = DSN()
	if err != nil {
		log.Fatal("Database configuration error: " + err.Error())
	}

	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	err = db.AutoMigrate(&models.Item{}, &models.Order{})
	if err != nil {
		log.Fatal("Migration failed: " + err.Error())
	}

	return
}
