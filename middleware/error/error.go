package err

import (
	"fmt"
	"khajuraho/backend/config"
	"khajuraho/backend/utils"

	"github.com/gofiber/fiber/v2"
)

// This Handler is a middleware that handles errors in the application. It
// prints the stack trace when in development mode and returns a JSON response
// with the error message. When in production mode, it simply returns a JSON
// response with a generic error message.
func Handler(c *fiber.Ctx, err error) error {
	env := config.AppConfig.Env
	errorMessage := utils.ServerSideErrorMessage

	if env == "development" {
		errorMessage = err.Error()
		fmt.Printf("Error %s", errorMessage)
	}

	return utils.ServerError(c, utils.ErrorMessage, []string{errorMessage}, nil)
}
