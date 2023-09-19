package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luissena/golang-crud/schemas"
)

func CreatePost(ctx *gin.Context) {
	request := CreatePostRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	post := schemas.Post{
		Author: request.Author,
		Title:  request.Title,
		Body:   request.Body,
	}

	if err := db.Create(&post).Error; err != nil {
		logger.Errorf("error creating post: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating post on database")
		return
	}

	sendSucess(ctx, "create-post", 201, post)

}
