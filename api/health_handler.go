package api

import (
	"khajuraho/backend/utils"

	"github.com/gofiber/fiber/v2"
)

// health
// @Summary Health Check
// @Description This endpoint checks api health.
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} dto.SendSuccess
// @Failure 401 {object} dto.SendError
// @Failure 429 {object} dto.SendError
// @Router /health [get]
func healthCheckHandler(c *fiber.Ctx) error {
	//TODO: check database and other services health and return the status in response.
	return utils.OK(c, nil, "Server is running")
}
