package router

import "github.com/gin-gonic/gin"

func Inititialzie() {
	router := gin.Default()
	startRoutes(router)
	router.Run()
}
