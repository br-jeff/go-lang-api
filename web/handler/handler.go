package handler

import (
	"gorm.io/gorm"

	"github.com/br-jeff/go-lang-api/internal/config"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandlers() {
	logger = config.GetLogger("Init")
	db = config.GetPG()
}
