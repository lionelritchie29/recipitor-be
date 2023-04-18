package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lionelritchie29/recipitor-be/database"
	"github.com/lionelritchie29/recipitor-be/models"
	"github.com/lionelritchie29/recipitor-be/routes"
)

func main() {
	godotenv.Load()
	initDb()
	router := gin.Default()
	port := os.Getenv("APP_PORT")

	apiRoute := router.Group("api")
	{
		routes.TestRoutes(apiRoute)
		routes.ItemRoutes(apiRoute)
	}
	
	router.Run(port)
}

func initDb() {
  db, err := database.GetConnection()

	if err != nil {
		log.Fatal(err)
		return
	}

	db.AutoMigrate(models.Item{})
	database.Seed(db)
}