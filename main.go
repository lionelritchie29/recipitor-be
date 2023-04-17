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
	
	router.GET("/ping", func (c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})

	routes.ItemRoutes(router)
	
	router.Run(port)
}