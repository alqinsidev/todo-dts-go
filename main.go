package main

import (
	"TODO-DIGITALENT/config"
	"TODO-DIGITALENT/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT must be set")
	}
	config.ConnectDB()
	server := gin.Default()
	routes.Route(server)
	server.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
