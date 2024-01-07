package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func productRoutes(prefixRouter *gin.RouterGroup) {
	prefixRouter.GET("/product", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "this is Product get",
		})
	})
}
