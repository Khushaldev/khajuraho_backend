package api

import (
	"khajuraho/backend/dto"
	"khajuraho/backend/service"
	"khajuraho/backend/utils"

	"github.com/gofiber/fiber/v2"
)

func googleAuthHandler(c *fiber.Ctx) error {
	var req dto.GoogleLoginRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, utils.BadRequestMessage, err) // TODO: need to change the error type
	}

	errors := utils.AppValidator(&req, c)
	if errors != nil {
		return utils.Unprocessable(c, "Data validation failed", *errors)
	}

	user, tokens, err := service.LoginWithGoogle(req)
	if err != nil {
		return utils.ServerError(c, "Login failed", err)

	}

	response := dto.GoogleLoginResponse{
		User: *user,
		Tokens: dto.AuthTokenResponse{
			AccessToken:  tokens.AccessToken.Token,
			RefreshToken: tokens.RefreshToken.Token,
		},
	}

	return utils.OK(c, "Login successful", response)
}

func logoutHandler(c *fiber.Ctx) error {

	var req dto.LogoutRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, utils.BadRequestMessage, err) // TODO: need to change the error type
	}

	errors := utils.AppValidator(&req, c)
	if errors != nil {
		return utils.Unprocessable(c, "Data validation failed", *errors)
	}

	if err := service.LogoutService(req.RefreshToken); err != nil {
		return utils.ServerError(c, "Logout failed", err)
	}

	return utils.OK(c, "Logout successful", nil)
}
