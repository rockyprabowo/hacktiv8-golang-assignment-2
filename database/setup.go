package database

import (
	"fmt"
	"rockyprabowo/assignment-2/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConfigMap map[string]string

func Dsn() string {
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

func Init() (db *gorm.DB) {
	var err error
	dsn := Dsn()
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&models.Item{}, &models.Order{})
	if err != nil {
		panic("Migration failed: " + err.Error())
	}

	return
}
