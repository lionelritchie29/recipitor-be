package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lionelritchie29/recipitor-be/database"
	"github.com/lionelritchie29/recipitor-be/models"
	"github.com/lionelritchie29/recipitor-be/models/dto"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {}

func (auth Auth) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := database.GetConnection()

		var payload dto.RegisterDto

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"data": nil,
				"success": false,
			})
			return
		}


		existResult := db.Where("email", payload.Email).First(&models.User{})

		if (existResult.RowsAffected > 0) {
			c.JSON(201, gin.H{
				"message": "user already exist",
				"data": nil,
				"success": false,
			})
			return
		}

		passwordBytes, _ := bcrypt. GenerateFromPassword([]byte(payload.Password), 14)
		user :=  models.User {
			Email: payload.Email,
			Password: string(passwordBytes),
		}
		
		db.Create(&user)
		c.JSON(201, gin.H{
			"message": "user registered!",
			"data": user,
			"success": true,
		})
	}
}