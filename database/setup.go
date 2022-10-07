package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"rocky.my.id/git/h8-assignment-2/models"
)

// DSN returns the DSN of the database connection.
func DSN() string {
	config := Config()
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s",
		config["host"],
		config["user"],
		config["password"],
		config["databaseName"],
		config["sslMode"],
	)
}

// Init initialise the database connection.
func Init() (db *gorm.DB) {
	var err error
	dsn := DSN()
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
