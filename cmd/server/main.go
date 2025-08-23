package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/eralora/mini-market-inventory-app/internal/database"
	"github.com/eralora/mini-market-inventory-app/internal/handler"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	if err := database.InitDB(); err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Here is my APP is OK"))
	})

	r.Route("/api", func(api chi.Router) {
		api.Post("/products", handler.CreateProduct)
		api.Get("/products", handler.ListProducts)
		api.Post("/stock/in", handler.AddStockIn)
		api.Post("/stock/out", handler.AddStockOut)
		api.Get("/inventory", handler.GetInventory)
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
