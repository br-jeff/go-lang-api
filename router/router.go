package router

import (
	"github.com/gin-gonic/gin"

	"github.com/br-jeff/go-lang-api/handler"
)

func Initialize() {
	router := gin.Default()
	handler.InitHandler()
	startRoutes(router)
	router.Run()
}
