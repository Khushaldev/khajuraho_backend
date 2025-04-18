package main

import (
	"log"

	"khajuraho/backend/internal/database"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Khajuraho First endpoint")
	})

	log.Fatal(app.Listen(":3000"))
}
