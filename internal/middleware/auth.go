package middleware

import (
	"fmt"
	"log"
	"strings"

	"khajuraho/backend/internal/config"
	"khajuraho/backend/pkg/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(config.AppConfig.JWTSecret)

func isValidJWT(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	return err == nil && token.Valid
}

func JWTMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenStr == "" || !isValidJWT(tokenStr) {
			return utils.Unauthorized(c, "Access denied")
		}

		return c.Next()
	}
}

// ClientSecretMiddleware is a middleware that validates the Client-Secret header in the request.
// The Client-Secret header is required and must match the value set in the environment variable.
// If the header is missing or the values do not match, it returns 401 Unauthorized.
func ClientSecretMiddleware() fiber.Handler {
	secret := config.AppConfig.ClientSecret

	if secret == "" {
		log.Fatal("FATAL: Client secret is not set in environment.")
	}

	return func(c fiber.Ctx) error {
		provided := c.Get("Client-Secret")

		if provided == "" {
			return utils.Unauthorized(c, "Missing Client-Secret header")
		}

		if provided != secret {
			return utils.Unauthorized(c, "Invalid Client-Secret")
		}

		return c.Next()
	}
}
