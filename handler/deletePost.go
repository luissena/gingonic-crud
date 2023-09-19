package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luissena/golang-crud/schemas"
)

func DeletePost(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	post := schemas.Post{}
	if err := db.First(&post, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("post with id %s not found", id))
		return
	}
	if err := db.Delete(&post).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting post with id %s", id))
		return
	}

	sendSucess(ctx, "delete-post", 200, post)
}
