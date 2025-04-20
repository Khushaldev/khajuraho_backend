package router

import (
	handler "khajuraho/backend/internal/handler/auth"

	"github.com/gofiber/fiber/v3"
)

func CreateRoutes(app *fiber.App) {
	v1Group := app.Group("/api/v1")

	authGroup := v1Group.Group("/auth")
	authGroup.Get("/send-otp", handler.SendOTP)
}
