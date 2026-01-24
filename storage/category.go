package storage

import (
	"errors"
	"sync"

	"simple-cashier-api/models"
)

var (
	categoryList []models.Category
	categoryMu   sync.RWMutex
)

func init() {
	categoryList = []models.Category{
		{ID: 1, Name: "Makanan", Description: "Kategori untuk makanan"},
		{ID: 2, Name: "Minuman", Description: "Kategori untuk minuman"},
		{ID: 3, Name: "Sembako", Description: "Kategori untuk sembako"},
	}
}

func GetAllCategories() []models.Category {
	categoryMu.RLock()
	defer categoryMu.RUnlock()
	return categoryList
}

func GetCategoryByID(id int) (models.Category, error) {
	categoryMu.RLock()
	defer categoryMu.RUnlock()

	for _, c := range categoryList {
		if c.ID == id {
			return c, nil
		}
	}
	return models.Category{}, errors.New("Category belum ada")
}

func AddCategory(category models.Category) models.Category {
	categoryMu.Lock()
	defer categoryMu.Unlock()

	category.ID = len(categoryList) + 1
	categoryList = append(categoryList, category)
	return category
}

func UpdateCategory(id int, updatedCategory models.Category) (models.Category, error) {
	categoryMu.Lock()
	defer categoryMu.Unlock()

	for i := range categoryList {
		if categoryList[i].ID == id {
			updatedCategory.ID = id
			categoryList[i] = updatedCategory
			return updatedCategory, nil
		}
	}
	return models.Category{}, errors.New("Category belum ada")
}

func DeleteCategory(id int) error {
	categoryMu.Lock()
	defer categoryMu.Unlock()

	for i, c := range categoryList {
		if c.ID == id {
			categoryList = append(categoryList[:i], categoryList[i+1:]...)
			return nil
		}
	}
	return errors.New("Category belum ada")
}
