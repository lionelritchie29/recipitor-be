package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lionelritchie29/recipitor-be/controllers"
)

func ItemRoutes(routerGroup *gin.RouterGroup) {
	itemRoutes := routerGroup.Group("items") 
	{
		itemRoutes.GET("/", controllers.Item{}.GetItems())
	}
}