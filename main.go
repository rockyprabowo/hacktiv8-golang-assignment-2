package main

import (
	"github.com/joho/godotenv"
	"log"

	"rockyprabowo/assignment-2/api/orders"
	"rockyprabowo/assignment-2/database"
	"rockyprabowo/assignment-2/routers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router := routers.NewDefaultRouter()
	db := database.Init()

	order_module.SetupDefault(router, db)

	err := router.Run(":5001")
	if err != nil {
		log.Fatal(err.Error())
	}
}
