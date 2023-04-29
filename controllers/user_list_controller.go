package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lionelritchie29/recipitor-be/database"
	"github.com/lionelritchie29/recipitor-be/models"
	"github.com/lionelritchie29/recipitor-be/models/dto"
)

type UserList struct {}

func (userList UserList) Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := database.GetConnection()
		userIdStr := c.Param("userId")

		var payload dto.CreateUserListDto

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"data": nil,
				"success": false,
			})
			return
		}

		userId, _ := strconv.Atoi(userIdStr)
		list := models.List {
			Description: payload.Description,
			UserId: userId,
		}

		db.Create(&list)

		for _, item := range payload.Items {
			list := models.ListItem {
				ItemId: item.Id,
				Amount: item.Amount,
				Quantity: item.Quantity,
				ListId: list.ID,
			}
	
			db.Create(&list)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Add new list",
			"data": list,
			"success": true,
		})
	}
}