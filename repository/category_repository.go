package repository

import (
	db "khajuraho/backend/database"
	"khajuraho/backend/model"
)

func FindCategories() (*[]model.Category, error) {
	var categories []model.Category
	result := db.Instance.Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return &categories, nil
}

func FindSubCategoriesByCategoryId(categoryId string) (*[]model.SubCategory, error) {
	var subCategories []model.SubCategory
	result := db.Instance.Where("category_id = ?", categoryId).Find(&subCategories)

	if result.Error != nil {
		return nil, result.Error
	}

	return &subCategories, nil
}

func FindSubCategoryById() error {
	return nil
}
