package main

import (
	"TODO-DIGITALENT/config"
	"TODO-DIGITALENT/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	server := gin.Default()
	routes.Route(server)
	server.Run(":" + config.EnvPort()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
