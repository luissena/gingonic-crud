package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luissena/golang-crud/schemas"
)

func ShowPostHandler(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	post := schemas.Post{}
	if err := db.First(&post, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("post with id %s not found", id))
		return
	}
	sendSucess(ctx, "show-post", 200, post)
}
