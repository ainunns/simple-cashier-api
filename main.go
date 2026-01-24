package main

import (
	"fmt"
	"log"
	"net/http"

	"simple-cashier-api/handlers"
)

func main() {
	http.HandleFunc("/api/produk", handlers.ProdukHandler)
	http.HandleFunc("/api/produk/", handlers.ProdukHandler)
	http.HandleFunc("/api/categories", handlers.CategoryHandler)
	http.HandleFunc("/api/categories/", handlers.CategoryHandler)
	http.HandleFunc("/health", handlers.HealthCheck)

	fmt.Println("Server running di localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Gagal running server:", err)
	}
}
