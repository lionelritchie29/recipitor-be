package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lionelritchie29/recipitor-be/controllers"
)

func AuthRoutes(routerGroup *gin.RouterGroup) {
	authRoutes := routerGroup.Group("auth") 
	{
		authRoutes.POST("register", controllers.Auth{}.Register())	
		authRoutes.POST("login", controllers.Auth{}.Login())	
	}
}