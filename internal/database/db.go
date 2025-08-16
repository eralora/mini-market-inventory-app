package database

import (
	"log"
    "github.com/eralora/mini-market-inventory-app\internal\model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("market.db"), &gorm.Config{})
	if err != nill {
		log.Fatal("failed to connect database: ", err)
	}
	// Auto-migrate tables
	DB.AutoMigrate(&model.Product{}, &model.StockEntry{})
}
