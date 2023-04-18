package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lionelritchie29/recipitor-be/routes"
)

func main() {
	godotenv.Load()
	router := gin.Default()
	port := os.Getenv("APP_PORT")

	apiRoute := router.Group("api")
	{
		routes.TestRoutes(apiRoute)
		routes.ItemRoutes(apiRoute)
	}
	
	

	
	router.Run(port)
}