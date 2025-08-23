package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/eralora/mini-market-inventory-app/internal/database"
	"github.com/eralora/mini-market-inventory-app/internal/model"
)

type stockReq struct {
	ProductID uint    `json:"product_id"`
	Quantity  float64 `json:"quantity"`
	Note      string  `json:"note"`
}

func addStock(w http.ResponseWriter, r *http.Request, typ string) {
	var req stockReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		badRequest(w, "invalid JSON")
		return
	}
	if req.ProductID == 0 || req.Quantity <= 0 {
		badRequest(w, "product_id and positive quantity required")
		return
	}

	// ensure product exists
	var p model.Product
	if err := database.DB.First(&p, req.ProductID).Error; err != nil {
		badRequest(w, "product not found")
		return
	}

	entry := model.StockEntry{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Type:      strings.ToUpper(typ), // "IN" or "OUT"
		Note:      req.Note,
	}
	if err := database.DB.Create(&entry).Error; err != nil {
		serverErr(w, err)
		return
	}

	// Reload with related product details
	if err := database.DB.Preload("Product").First(&entry, entry.ID).Error; err != nil {
		serverErr(w, err)
		return
	}

	writeJson(w, http.StatusCreated, entry)
}

func AddStockIn(w http.ResponseWriter, r *http.Request)  { addStock(w, r, "IN") }
func AddStockOut(w http.ResponseWriter, r *http.Request) { addStock(w, r, "OUT") }
