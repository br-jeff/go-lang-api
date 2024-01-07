package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthRoute(router *gin.Engine) {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Server is runing",
		})
	})
}
