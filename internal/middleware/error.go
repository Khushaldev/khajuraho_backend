package middleware

import (
	"fmt"
	"khajuraho/backend/internal/config"
	"runtime/debug"

	"github.com/gofiber/fiber/v3"
)

// ErrorHandler is a middleware that handles errors in the application. It
// prints the stack trace when in development mode and returns a JSON response
// with the error message. When in production mode, it simply returns a JSON
// response with a generic error message.
func ErrorHandler(c fiber.Ctx, err error) error {
	env := config.AppConfig.Env

	if env == "development" {
		// Print stack trace
		fmt.Printf("%s", debug.Stack())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Internal Server Error",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"success": false,
		"message": "Something went wrong",
		"errors":  "Server Error, please try again later",
	})
}
