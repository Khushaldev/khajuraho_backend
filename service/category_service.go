package service

import (
	"khajuraho/backend/dto"
	"khajuraho/backend/model"
	"khajuraho/backend/repository"
	"khajuraho/backend/service/cache"

	"github.com/google/uuid"
)

func BuildCategoryTree(categories []model.Category, parentID *uuid.UUID) []dto.CategoryNode {
	var childNodes []dto.CategoryNode

	for _, category := range categories {
		isDirectChild := (category.ParentID == nil && parentID == nil) || (category.ParentID != nil && parentID != nil && *category.ParentID == *parentID)
		if isDirectChild {
			childNode := dto.CategoryNode{
				ID:       category.ID.String(),
				Name:     category.Name,
				Slug:     category.Slug,
				Icon:     category.Icon,
				Children: BuildCategoryTree(categories, &category.ID),
			}
			childNodes = append(childNodes, childNode)
		}
	}

	return childNodes
}

func GetCategoryNode() ([]dto.CategoryNode, error) {
	var categories []model.Category
	err := cache.GetOrSetCache(cache.CategoriesKey, &categories, repository.FindCategories)
	if err != nil {
		return nil, err
	}

	tree := BuildCategoryTree(categories, nil)
	return tree, nil
}
