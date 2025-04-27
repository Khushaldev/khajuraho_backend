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
		return utils.BadRequest(c, "Invalid request format", []string{"Invalid request"}, err)
	}

	user, tokens, err := service.LoginWithGoogle(req)
	if err != nil {
		return utils.ServerError(c, "Login failed", []string{"Something went wrong"}, err)

	}

	response := dto.GoogleLoginResponse{
		User: *user,
		Tokens: dto.AuthTokenResponse{
			AccessToken:  tokens.AccessToken.Token,
			RefreshToken: tokens.RefreshToken.Token,
		},
	}

	return utils.OK(c, response, "Login successful")
}

func logoutHandler(c *fiber.Ctx) error {

	var req dto.LogoutRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "Invalid request format", []string{"Invalid request"}, err)
	}

	if err := service.LogoutService(req.RefreshToken); err != nil {
		return utils.ServerError(c, "Logout failed", []string{"Something went wrong"}, err)
	}

	return utils.OK(c, nil, "Logout successful")
}
