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
		v1.GET("/posts", handler.ListPostsHandler)
		v1.GET("/posts/:id", handler.ShowPostHandler)
		v1.POST("/posts", handler.CreatePostHandler)
		v1.DELETE("/posts", handler.DeletePostHandler)
		v1.PUT("/posts", handler.UpdatePostHandler)
	}

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
}
