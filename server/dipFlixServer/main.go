package main

import (
	"fmt"

	"github.com/khdip/dip-flix/server/dipFlixServer/database"
	"github.com/khdip/dip-flix/server/dipFlixServer/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	router := gin.Default()
	var client *mongo.Client = database.Connect()

	routes.SetupUnprotectedRoutes(router, client)
	routes.SetupProtectedRoutes(router, client)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}
