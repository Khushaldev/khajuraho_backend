package api

import (
	"khajuraho/backend/service"
	"khajuraho/backend/utils"

	"github.com/gofiber/fiber/v2"
)

func getCategoryNodeHandler(c *fiber.Ctx) error {
	categories, err := service.GetCategoryNode()
	if err != nil || categories == nil {
		return utils.ServerError(c, "Failed to fetch categories ", err)

	}
	// type FlatAddressResponse struct {
	// 	Line1   string `json:"line1"`
	// 	Line2   string `json:"line2"`
	// 	Area    string `json:"area"`
	// 	City    string `json:"city"`
	// 	State   string `json:"state"`
	// 	Country string `json:"country"`
	// }

	// var address model.Address
	// db.Instance.Preload("Area.City.State.Country").First(&address)

	// response := FlatAddressResponse{
	// 	Line1:   address.Line1,
	// 	Line2:   address.Line2,
	// 	Area:    address.Area.Name,
	// 	City:    address.Area.City.Name,
	// 	State:   address.Area.City.State.Name,
	// 	Country: address.Area.City.State.Country.Name,
	// }

	// data := map[string]any{
	// 	"categories": categories,
	// 	"address":    response,
	// }

	return utils.OK(c, "Categories fetched successfully", categories)
}
