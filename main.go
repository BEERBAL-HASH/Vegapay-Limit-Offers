package main

import (
	"log"

	"github.com/BEERBAL-HASH/Vegapay-Limit-Offer/database"
	"github.com/BEERBAL-HASH/Vegapay-Limit-Offer/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//creating instance of gin router, default() method gives you logging as well
	router := gin.Default()

	//connecting database to the server
	database.ConnectDB()

	//setup routes
	routes.SetupRoutes(router)

	//running server instance on localhost port 80
	log.Fatal(router.Run("localhost:8080"))
}
