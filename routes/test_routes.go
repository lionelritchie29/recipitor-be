package routes

import "github.com/gin-gonic/gin"

func TestRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/ping", func (c *gin.Context) {
		c.String(200, "pong")
	})
}