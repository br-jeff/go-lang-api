package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/br-jeff/go-lang-api/schemas"
)

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func (r *CreateProductRequest) validate() error {
	if r.Name == "" && r.Price == 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Price == 0 {
		return errParamIsRequired("name", "number")
	}
	return nil
}

func CreateProductHandler(ctx *gin.Context) {
	request := CreateProductRequest{}

	ctx.BindJSON(request)

	if err := request.validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
	}

	product := schemas.Product{
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
