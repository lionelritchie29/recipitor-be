package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
			c.JSON(200, gin.H{
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

func (auth Auth) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := database.GetConnection()

		var payload dto.LoginDto

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
				"data": nil,
				"success": false,
			})
			return
		}
		
		var user models.User
		existUser := db.Where("email", payload.Email).First(&user)

		if (existUser.RowsAffected == 0 ){
			c.JSON(200, gin.H{
				"message": "user does not exists",
				"data": nil,
				"success": false,
			})
			return
		}

		hashMatchErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
		if hashMatchErr != nil {
			c.JSON(200, gin.H{
				"message": "Wrong password",
				"data": nil,
				"success": false,
			})
			return
		}

		key := os.Getenv("JWT_kEY")
		t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.ID,
			"email": user.Email,
		})
		token, _ := t.SignedString([]byte(key))

		c.JSON(200, gin.H {
			"message": "Login success!",
			"data": token,
			"success": true,
		})		
	}
}