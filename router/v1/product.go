package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func productRoutes(prefixRouter *gin.RouterGroup) {
	prefixRouter.GET("/product", getProductHandler)
}

func getProductHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "this is Product get",
	})
}
