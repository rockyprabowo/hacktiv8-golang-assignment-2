package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
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
	// Setup server listen and serve address
	serveAddress := setupServerAddr()

	// Swagger documentations runtime values
	docs.SwaggerInfo.Host = options.Default(os.Getenv("APP_STATIC_URL"), serveAddress)
	docs.SwaggerInfo.BasePath = "/"

	// Setup Gin router
	setupGinLogging()
	router := routers.NewDefaultRouter()

	// Initialise database connection
	db := database.Init()

	// Load Order module
	order_module.SetupDefault(router, db)

	// Swagger documentations route setup
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the server
	err := router.Run(serveAddress)
	if err != nil {
		log.Fatal(err.Error())
	}
}
