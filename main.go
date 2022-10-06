package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"rockyprabowo/assignment-2/helpers/options"

	"rockyprabowo/assignment-2/api/orders"
	"rockyprabowo/assignment-2/database"
	"rockyprabowo/assignment-2/routers"
)

func setupServerAddr() string {
	host := options.Default(os.Getenv("APP_HOST"), "localhost")
	port := options.Default(os.Getenv("APP_PORT"), "5001")
	return fmt.Sprintf("%s:%s", host, port)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	serveAddress := setupServerAddr()
	router := routers.NewDefaultRouter()
	db := database.Init()

	order_module.SetupDefault(router, db)

	err := router.Run(serveAddress)
	if err != nil {
		log.Fatal(err.Error())
	}
}
