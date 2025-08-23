package database

import (
	"github.com/eralora/mini-market-inventory-app/internal/model"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error

	DB, err = gorm.Open(sqlite.Open("market.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto-migrate creates tables if not present
	if err := DB.AutoMigrate(&model.Product{}, &model.StockEntry{}); err != nil {
		return err
	}

	return nil
}
