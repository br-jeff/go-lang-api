package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func startRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Server is runing",
			})
		})
	}

}
