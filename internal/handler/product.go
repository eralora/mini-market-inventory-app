package handler

import (
	"encoding/json"
	"net/http"

	"github.com/eralora/mini-market-inventory-app/internal/database"
	"github.com/eralora/mini-market-inventory-app/internal/model"
)

type createProductReq struct {
	Name  string  `json:"name"`
	Unit  string  `json:"unit"`
	Price float64 `json:"price"`
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req createProductReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		badRequest(w, "invalid JSON")
		return
	}
	if req.Name == "" || req.Unit == "" {
		badRequest(w, "name and unit are required")
		return
	}
	p := model.Product{Name: req.Name, Unit: req.Unit, Price: req.Price}
	if err := database.DB.Create(&p).Error; err != nil {
		serverErr(w, err)
		return
	}
	writeJson(w, http.StatusCreated, p)
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	var products []model.Product
	if err := database.DB.Order("id asc").Find(&products).Error; err != nil {
		serverErr(w, err)
		return
	}
	writeJson(w, http.StatusOK, products)
}
