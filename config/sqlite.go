package config

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/br-jeff/go-lang-api/schemas"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	sqlitePath := "./db/main.db"
	_, err := os.Stat(sqlitePath)

	if os.IsNotExist(err) {
		logger.Info("Database file not found")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(sqlitePath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Erro when try to open sqlite: %v", err)
	}

	err = db.AutoMigrate(&schemas.Product{})

	if err != nil {
		logger.Errorf("Sqlite automagration error %v", err)
		return nil, err
	}

	return db, nil
}
