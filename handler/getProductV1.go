package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "this is Product get",
	})
}
