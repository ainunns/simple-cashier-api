package storage

import (
	"errors"
	"sync"

	"simple-cashier-api/models"
)

var (
	produkList []models.Produk
	mu         sync.RWMutex
)

func init() {
	produkList = []models.Produk{
		{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
		{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
		{ID: 3, Nama: "kecap", Harga: 12000, Stok: 20},
	}
}

func GetAllProduk() []models.Produk {
	mu.RLock()
	defer mu.RUnlock()
	return produkList
}

func GetProdukByID(id int) (models.Produk, error) {
	mu.RLock()
	defer mu.RUnlock()

	for _, p := range produkList {
		if p.ID == id {
			return p, nil
		}
	}
	return models.Produk{}, errors.New("Produk belum ada")
}

func AddProduk(produk models.Produk) models.Produk {
	mu.Lock()
	defer mu.Unlock()

	produk.ID = len(produkList) + 1
	produkList = append(produkList, produk)
	return produk
}

func UpdateProduk(id int, updatedProduk models.Produk) (models.Produk, error) {
	mu.Lock()
	defer mu.Unlock()

	for i := range produkList {
		if produkList[i].ID == id {
			updatedProduk.ID = id
			produkList[i] = updatedProduk
			return updatedProduk, nil
		}
	}
	return models.Produk{}, errors.New("Produk belum ada")
}

func DeleteProduk(id int) error {
	mu.Lock()
	defer mu.Unlock()

	for i, p := range produkList {
		if p.ID == id {
			produkList = append(produkList[:i], produkList[i+1:]...)
			return nil
		}
	}
	return errors.New("Produk belum ada")
}
