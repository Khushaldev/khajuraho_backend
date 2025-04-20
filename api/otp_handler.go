package api

import (
	"khajuraho/backend/services"
	"khajuraho/backend/utils"

	"github.com/gofiber/fiber/v2"
)

// sendOTP
// @Summary Send OTP
// @Description This endpoint sends OTP.
// @Tags OTP
// @Accept json
// @Produce json
// @Success 204 "OTP successfully sent"
// @Router /api/v1/auth/send-otp [get]
func sendOTP(c *fiber.Ctx) error {
	// TODO: Send OTP, remove mock response

	err := services.SendOTP()
	if err != nil {
		return utils.ServerError(c)
	}

	return utils.Success(c, "", "")
}
