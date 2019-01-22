package route

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
)

func Route() *gin.Engine {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("../frontend/dist", false)))
	api := router.Group("/api")
	{
		apia(api)
	}
	return router
}

func apia(router *gin.RouterGroup) {
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"test": "hello world"})
	})
}
