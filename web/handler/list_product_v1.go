package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/br-jeff/go-lang-api/internal/models"
)

func ListProductHandler(ctx *gin.Context) {
	products := models.Product{}
	db.Find(&products)
	if products.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Error not found"})
		return
	}
	ctx.ShouldBindJSON(&products)

	ctx.JSON(http.StatusOK, products)
}
