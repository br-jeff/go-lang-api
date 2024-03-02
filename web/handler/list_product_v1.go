package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/br-jeff/go-lang-api/internal/models"
)

func ListProductHandler(ctx *gin.Context) {
	products := models.Product{}
	db.Find(&products)
	ctx.ShouldBindJSON(&products)
	ctx.JSON(http.StatusOK, products)
}
