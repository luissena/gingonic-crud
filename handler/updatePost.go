package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luissena/golang-crud/schemas"
)

func UpdatePost(ctx *gin.Context) {
	request := UpdatePostRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
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
	if request.Author != "" {
		post.Author = request.Author
	}
	if request.Title != "" {
		post.Title = request.Title
	}
	if request.Body != "" {
		post.Body = request.Body
	}
	if err := db.Save(&post).Error; err != nil {
		logger.Errorf("error updating post with id %s: %v", id, err.Error())
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error updating post with id %s", id))
		return
	}
	sendSucess(ctx, "update-post", 200, post)
}
