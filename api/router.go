package api

import (
	"khajuraho/backend/middleware/auth"
	"khajuraho/backend/models"
	"khajuraho/backend/utils"
	"time"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

func SetupRoutes(app *fiber.App) {
	// API group with logging and client secret middleware
	api := app.Group("/api", logger.New(), auth.RequireClientSecret())

	app.Get("/docs/*", swagger.HandlerDefault)

	apiv1 := api.Group("/v1")

	app.Use(logger.New(
		logger.Config{
			Format:     "${cyan}[${time}] ${green}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
			TimeFormat: "02-Jan-2006",
			TimeZone:   "Asia/Kolkata",
		},
	))

	// Auth route
	auth := apiv1.Group("/auth")

	auth.Get("/send-otp", sendOTP)

	// TODO: health check add
}

// sendOTP
// @Summary Send OTP
// @Description This endpoint sends OTP.
// @Tags OTP
// @Accept json
// @Produce json
// @Success 204 "OTP successfully sent"
// @Router /api/v1/auth/send-otp [get]
func sendOTP(c *fiber.Ctx) error {
	user := models.User{
		ID:        uuid.New(),
		Name:      "Khushal",
		Email:     "jW2yS@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return utils.Success(c, user, utils.SuccessMessage)
}
