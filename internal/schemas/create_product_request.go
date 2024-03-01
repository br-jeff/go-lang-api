package schemas

import (
	"fmt"
)

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func (r *CreateProductRequest) Validate() error {
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
