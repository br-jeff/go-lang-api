package router

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/br-jeff/go-lang-api/router/v1"
)

func startRoutes(router *gin.Engine) {
	healthRoute(router)
	v1Router := router.Group("api/v1")
	{
		v1.GetRoutes(v1Router)
	}

}
