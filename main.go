package main

import (
	"log"
	"rockyprabowo/assignment-2/api/orders"
	"rockyprabowo/assignment-2/database"
	"rockyprabowo/assignment-2/routers"
)

func main() {
	router := routers.NewDefaultRouter()
	db := database.Init()

	order_module.SetupDefault(router, db)

	err := router.Run(":5001")
	if err != nil {
		log.Fatal(err.Error())
	}
}
