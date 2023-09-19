package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSucess(ctx *gin.Context, op string, code int, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"code":    code,
		"message": fmt.Sprintf("Operation from handler: %s sucessfull", op),
		"data":    data,
	})
}
