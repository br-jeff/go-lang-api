package router

import (
	"github.com/gin-gonic/gin"

	"github.com/br-jeff/go-lang-api/web/handler"
)

func Initialize() {
	router := gin.Default()
	handler.InitHandlers()
	bootstrap(router)
	router.Run()
}
