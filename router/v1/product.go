package apiv1

import (
	"github.com/gin-gonic/gin"

	"github.com/br-jeff/go-lang-api/handler"
)

func productRoutes(rg *gin.RouterGroup) {
	rg.GET("/product", handler.GetProductHandler)
	rg.POST("/product", handler.CreateProductHandler)
}
