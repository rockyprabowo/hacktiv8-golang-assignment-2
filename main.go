package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"os"
	"rocky.my.id/git/h8-assignment-2/api/orders"
	"rocky.my.id/git/h8-assignment-2/database"
	"rocky.my.id/git/h8-assignment-2/docs"
	"rocky.my.id/git/h8-assignment-2/helpers/options"
	"rocky.my.id/git/h8-assignment-2/routers"
)

// @title       Hacktiv8 Golang Assignment 2 - Orders
// @version     1.0
// @description This is a REST API for a custom shop ordering system.
//
// @contact.name  API Support
// @contact.url   https://rocky.my.id/
// @contact.email rocky@lazycats.id
func main() {
	// Loads the environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("WARN: Couldn't load .env file. Checking system environment variables.")
		if err := CheckEnvVars(); err != nil {
			log.Println(err.Error())
			log.Fatal("Couldn't configure the application from every environment variables!")
		}
	}

	// Setup server listen and serve address
	serveAddress := setupServerAddr()

	// Swagger documentations runtime values
	docs.SwaggerInfo.Host = options.Default(os.Getenv("APP_STATIC_URL"), serveAddress)
	docs.SwaggerInfo.BasePath = "/"

	// Setup router and initialise database connection
	router := routers.NewDefaultRouter()
	db := database.Init()

	// Load Order module
	order_module.SetupDefault(router, db)

	// Swagger documentations route setup
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the server
	err := router.Run(serveAddress)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// setupServerAddr builds the address for Gin to listen and serve the application to.
func setupServerAddr() string {
	host := options.Default(os.Getenv("APP_HOST"), "localhost")
	port := options.Default(os.Getenv("APP_PORT"), "5001")
	return fmt.Sprintf("%s:%s", host, port)
}
