package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luissena/golang-crud/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		v1.GET("/posts", handler.ListPosts)
	}

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
}
