package handler

import "github.com/gin-gonic/gin"

func ListPosts(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "ListPosts",
	})
}
