package err

import (
	"fmt"
	"khajuraho/backend/config"
	"khajuraho/backend/utils"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
)

// This Handler is a middleware that handles errors in the application. It
// prints the stack trace when in development mode and returns a JSON response
// with the error message. When in production mode, it simply returns a JSON
// response with a generic error message.
func Handler(c *fiber.Ctx, err error) error {
	env := config.AppConfig.Env

	if env == "development" {
		// Print stack trace
		fmt.Printf("%s", debug.Stack())

		return utils.ServerError(c, utils.ErrorMessage, []string{err.Error()})
	}

	return utils.ServerError(c, utils.ErrorMessage, []string{utils.ServerSideErrorMessage})
}
