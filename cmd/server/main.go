package main

import (
	"log"
	"net/http"
	"github.com/eralora/mini-market-inventory-app\internal\database"
)

func main() {

	if err := database.InitDB(); err != nil {
		log.Fatal(err)
	}


	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Here is my APP"))
	})

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
