package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lionelritchie29/recipitor-be/database"
	"github.com/lionelritchie29/recipitor-be/models"
	"github.com/lionelritchie29/recipitor-be/models/dto"
)

type Item struct {}

func (item Item) GetItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := database.GetConnection()
		var items []models.Item

		db.Find(&items)

		c.JSON(200, gin.H{
			"message": "get items",
			"data": items,
			"success": true,
		})
	}
}

func (item Item) GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		db, _ := database.GetConnection()
		var item models.Item

		db.First(&item, id)

		c.JSON(200, gin.H{
			"message": "get single item",
			"data": item,
			"success": true,
		})
	}
}

func (item Item) AddItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := database.GetConnection()
		
		var payload dto.CreateItemDto
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"data": nil,
				"success": false,
			})
			return
		}
		
		newItem := models.Item {
			Name: payload.Name,
			Image: payload.Image,
			Description: payload.Description,
		}

		db.Create(&newItem)

		c.JSON(201, gin.H{
			"message": "create new item",
			"data": newItem,
			"success": true,
		})
	}
}

func (item Item) UpdateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := database.GetConnection()
		id := c.Param("id")
		
		var payload dto.UpdateItemDto
		var item models.Item

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"data": nil,
				"success": false,
			})
			return
		}

		result := db.First(&item, id)

		if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Item not exist",
				"data": nil,
				"success": false,
			})
			return
		}

		if (payload.Name != "") {
			item.Name = payload.Name
		}

		if (payload.Description != "") {
			item.Description = payload.Description
		}

		if (payload.Image != "") {
			item.Image = payload.Image
		}

		db.Save(&item)

		c.JSON(201, gin.H{
			"message": "update new item",
			"data": item,
			"success": true,
		})
	}
}