package router

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/br-jeff/go-lang-api/web/router/v1"
)

func bootstrap(router *gin.Engine) {
	healthRoute(router)
	v1Router := router.Group("api/v1")
	v1.BootstrapApiV1(v1Router)
}
