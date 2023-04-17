package routes

import "github.com/gin-gonic/gin"

func ItemRoutes(router *gin.Engine) {
	itemRoutes := router.Group("items") 
	{
		itemRoutes.GET("/", func (c *gin.Context) {
			c.JSON(200, gin.H {
				"message": "get items",
			})
		})
	}
}