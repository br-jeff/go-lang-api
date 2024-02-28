package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/br-jeff/go-lang-api/internal/models"
)

func InitializePG() (*gorm.DB, error) {
	logger := GetLogger("DBPG")

	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  "host=postgres user=dev password=dev dbname=dev port=5432 sslmode=disable",
			PreferSimpleProtocol: true,
		}), &gorm.Config{})

	if err != nil {
		logger.Errorf("Erro when try to open PG: %v", err)
	}

	err = db.AutoMigrate(&models.Product{}, &models.User{})

	if err != nil {
		logger.Errorf("DB automagration error %v", err)
		return nil, err
	}

	return db, nil
}
