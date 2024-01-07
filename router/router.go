package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()
	startRoutes(router)
	router.Run()
}
