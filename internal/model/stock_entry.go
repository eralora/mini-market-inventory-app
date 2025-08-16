package model

import (
	"gorm.io/gorm"
)

type StockEntry struct {
	gorm.Model
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  float64 `gorm:"not null"`
	Type      string  `gorm:"not null"` // "IN" or "OUT"
	Note      string  // optional comment like "from wholesaler"
}
