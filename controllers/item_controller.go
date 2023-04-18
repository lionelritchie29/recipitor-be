package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lionelritchie29/recipitor-be/database"
	"github.com/lionelritchie29/recipitor-be/models"
)

type Item struct {}

func (c Item) GetItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := database.GetConnection()
		var items []models.Item

		db.Find(&items)

		c.JSON(200, gin.H{
			"message": "get items",
			"data": items,
		})
	}
}