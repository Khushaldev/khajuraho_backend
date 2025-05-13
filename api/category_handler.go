package api

import (
	"khajuraho/backend/service"
	"khajuraho/backend/utils"

	"github.com/gofiber/fiber/v2"
)

func getCategoryHierarchyHandler(c *fiber.Ctx) error {
	categories, err := service.GetCategoryHierarchy()
	if err != nil || categories == nil {
		return utils.ServerError(c, "Failed to fetch categories ", []string{"Something went wrong"}, err)

	}

	return utils.OK(c, categories, "Categories fetched successfully")
}
