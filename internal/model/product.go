package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model         // adds ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string  `gorm:"not null"`
	Unit       string  `gorm:"not null"` // e.g. pcs, kg, box
	Price      float64 `gorm:"not null"` // Purchase or sale price
}
