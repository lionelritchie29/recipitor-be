package routes

import "github.com/gin-gonic/gin"

func ItemRoutes(routerGroup *gin.RouterGroup) {
	itemRoutes := routerGroup.Group("items") 
	{
		itemRoutes.GET("/", func (c *gin.Context) {
			c.JSON(200, gin.H {
				"message": "get items",
			})
		})
	}
}