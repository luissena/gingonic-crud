package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luissena/golang-crud/schemas"
)

func ListPosts(ctx *gin.Context) {
	posts := []schemas.Post{}

	if err := db.Find(&posts).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error listing posts")
		return
	}
	sendSucess(ctx, "list-posts", 200, posts)
}
