package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lionelritchie29/recipitor-be/controllers"
)

func UserListRoutes(routerGroup *gin.RouterGroup) {
	userListRoutes := routerGroup.Group("user/:userId/items") 
	{
		userListRoutes.POST("", controllers.UserList{}.Add())
	}
}