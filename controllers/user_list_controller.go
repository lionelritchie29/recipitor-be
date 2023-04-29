package controllers

import (
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
			c.JSON(200, gin.H{
				"message": err.Error(),
				"data": nil,
				"success": false,
			})
			return
		}

		userId, _ := strconv.Atoi(userIdStr)
		list := models.List {
			Name: payload.Name,
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

		c.JSON(200, gin.H{
			"message": "Add new list",
			"data": list,
			"success": true,
		})
	}
}

func (userList UserList) GetByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := database.GetConnection()
		userIdStr := c.Param("userId")
		userId, _ := strconv.Atoi(userIdStr)

		var lists []models.List
		var userLists []dto.GetListDto

		db.Where("user_id", userId).Find(&lists)

		for _, list := range lists {
			var listItems []models.ListItem
			var listItemsDto []dto.GetListItemDto
			
			db.Where("list_id", list.ID).Preload("Item").Find(&listItems)

			for _, listItem := range listItems {
				listItemsDto = append(listItemsDto, dto.GetListItemDto{
					Id: listItem.Item.ID,
					Name: listItem.Item.Name,
					Image: listItem.Item.Image,
					Amount: listItem.Amount,
					Quantity: listItem.Quantity,
				})
			}
			
			dto := dto.GetListDto {
				Name: list.Name,
				Items: listItemsDto,
			}

			userLists = append(userLists, dto)
		}
		
		c.JSON(200, gin.H{
			"message": "Get all lists for current user",
			"data": userLists,
			"success": true,
		})
	}
}