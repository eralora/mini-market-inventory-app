package handler

import (
	"net/http"

	"github.com/eralora/mini-market-inventory-app/internal/database"
	"github.com/eralora/mini-market-inventory-app/internal/model"
)

type InventoryItem struct {
	ProductID  uint    `json:"product_id"`
	Name       string  `json:"name"`
	Unit       string  `json:"unit"`
	CurrentQty float64 `json:"current_qty"`
}

func GetInventory(w http.ResponseWriter, r *http.Request) {
	var products []model.Product
	if err := database.DB.Find(&products).Error; err != nil {
		serverErr(w, err)
		return
	}

	var rows []InventoryItem

	for _, p := range products {
		var totalIn float64
		var totalOut float64

		database.DB.
			Model(&model.StockEntry{}).
			Where("product_id = ? AND type = ?", p.ID, "IN").
			Select("SUM(quantity)").
			Scan(&totalIn)

		database.DB.
			Model(&model.StockEntry{}).
			Where("product_id = ? AND type = ?", p.ID, "OUT").
			Select("SUM(quantity)").
			Scan(&totalOut)

		rows = append(rows, InventoryItem{
			ProductID:  p.ID,
			Name:       p.Name,
			Unit:       p.Unit,
			CurrentQty: totalIn - totalOut,
		})
	}

	writeJson(w, http.StatusOK, rows)
}
