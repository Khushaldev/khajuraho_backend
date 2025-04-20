package middleware

import (
	"fmt"
	"log"
	"strings"

	"khajuraho/backend/internal/config"
	"khajuraho/backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// TODO: Need to improve JWT verification
func isValidJWT(tokenStr string) bool {
	var secretKey = []byte(config.AppConfig.JWTSecret)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	return err == nil && token.Valid
}

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
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
	clientKey := config.AppConfig.ClientKey
	clientSecret := config.AppConfig.ClientSecret

	if clientSecret == "" || clientKey == "" {
		log.Fatal("FATAL: Client secret is not set in environment.")
	}

	return func(c *fiber.Ctx) error {
		provided := c.Get(clientKey)

		if provided == "" {
			return utils.Unauthorized(c, "Missing Client-Secret header")
		}

		if provided != clientSecret {
			return utils.Unauthorized(c, "Invalid Client-Secret")
		}

		return c.Next()
	}
}
