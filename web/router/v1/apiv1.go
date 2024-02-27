package apiv1

import (
	"github.com/gin-gonic/gin"
)

func BootstrapApiV1(prefixRouter *gin.RouterGroup) {
	productRoutes(prefixRouter)
}
