package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/br-jeff/go-lang-api/internal/models"
	"github.com/br-jeff/go-lang-api/internal/schemas"
)

func CreateProductHandler(ctx *gin.Context) {
	request := schemas.CreateProductRequest{}

	ctx.BindJSON(request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
	}

	product := models.Product{
		Name:  request.Name,
		Price: request.Price,
	}

	logger.Infof("request recived", request)

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("erro on database: ", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error on database")
	}

	sendSucess(ctx, "Creating Product", product)
}
