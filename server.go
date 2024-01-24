package main

import (
	"myapp/config"
	"myapp/helpers"
	"myapp/http/routes"
	"myapp/migration"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func init() {
	helpers.LoadEnv()

	config.ConnectDB()

	migration.MigrateUp()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	r := gin.Default()
	routes.InitRoutes(r)

	r.Run("localhost:" + port)
}
