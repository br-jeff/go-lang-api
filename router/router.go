package router

import "github.com/gin-gonic/gin"

func Inititialzie() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is runing",
		})
	})

	router.Run()
}
