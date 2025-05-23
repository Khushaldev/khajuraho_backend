package api

import (
	"khajuraho/backend/middleware/auth"
	"khajuraho/backend/middleware/limiter"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	auth.Post("/google", googleAuthHandler)
	auth.Post("/logout", logoutHandler)

	category := apiv1.Group("/category")
	category.Get("/hierarchy", getCategoryNodeHandler)

	// Basic Health check with rate limiter
	app.Get("/health", healthCheckHandler, limiter.New())
}
