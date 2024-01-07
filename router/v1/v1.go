package apiv1

import (
	"github.com/gin-gonic/gin"
)

func GetRoutes(prefixRouter *gin.RouterGroup) {
	productRoutes(prefixRouter)
}
