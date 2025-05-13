package service

import (
	"fmt"
	"khajuraho/backend/dto"
	"khajuraho/backend/model"
	"khajuraho/backend/repository"
)

func GetCategories() (*[]model.Category, error) {
	categories, err := repository.FindCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func GetCategoryHierarchy() (*[]dto.CategoryHierarchy, error) {
	categories, err := repository.FindCategories()
	if err != nil {
		return nil, err
	}

	if categories == nil {
		return nil, fmt.Errorf("categories not found")
	}
	var hierarchy []dto.CategoryHierarchy

	for _, category := range *categories {
		subCategories, err := repository.FindSubCategoriesByCategoryId(category.ID.String())
		if err != nil || subCategories == nil {
			continue
		} else {
			var subCategoryHierarchy []dto.CategoryHierarchy
			for _, subCategory := range *subCategories {
				subCategoryHierarchy = append(subCategoryHierarchy, dto.CategoryHierarchy{
					ID:       subCategory.ID.String(),
					Name:     subCategory.Name,
					Slug:     subCategory.Slug,
					Icon:     subCategory.Icon,
					Children: []dto.CategoryHierarchy{},
				})
			}
			hierarchy = append(hierarchy, dto.CategoryHierarchy{
				ID:       category.ID.String(),
				Name:     category.Name,
				Slug:     category.Slug,
				Icon:     category.Icon,
				Children: subCategoryHierarchy,
			})
		}

	}

	return &hierarchy, nil
}
