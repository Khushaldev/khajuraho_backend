package api

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// API group with logging middleware
	api := app.Group("/api", logger.New())

	app.Get("/docs/*", swagger.HandlerDefault)

	apiv1 := api.Group("/v1")

	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${green}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Kolkata",
	}))

	// Auth route
	auth := apiv1.Group("/auth")
	auth.Get("/send-otp", sendOTP)

	// TODO: health check add
}
