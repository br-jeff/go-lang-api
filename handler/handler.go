package handler

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/br-jeff/go-lang-api/config"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSqlite()
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}
