package limiter

import (
	"khajuraho/backend/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func New() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:          20,
		Expiration:   30 * time.Second,
		LimitReached: limitReachedHandler,
	})
}

func limitReachedHandler(c *fiber.Ctx) error {
	return utils.TooManyRequests(c, utils.TooManyRequestsMessage, utils.ErrorCode{Code: utils.TooManyRequestsCode})
}
