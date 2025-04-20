package handler

import (
	"khajuraho/backend/utils"

	"github.com/gofiber/fiber/v2"
)

func SendOTP(c *fiber.Ctx) error {
	// TODO: Send OTP, remove mock response
	return utils.Success(c, "Data Received", "Successfully received data")
}
