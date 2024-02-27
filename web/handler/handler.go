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

func InitDatabase() {
	logger = config.GetLogger("InitDb")
	db = config.GetPG()
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}