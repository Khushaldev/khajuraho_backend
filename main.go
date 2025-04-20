package main

import (
	"log"
	"os"

	"khajuraho/backend/api"
	"khajuraho/backend/config"
	db "khajuraho/backend/database"
	"khajuraho/backend/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"

	_ "khajuraho/backend/docs"
)

func loggerFunction() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile | log.Lmicroseconds)
	log.SetOutput(os.Stdout)
}

// @title Backend Service API
// @version 1.0
// @description API documentation for Bonus Service
func main() {
	loggerFunction()
	config.LoadConfig()

	configuration := fiber.Config{
		ErrorHandler:            middleware.ErrorHandler,
		EnableTrustedProxyCheck: true,
	}

	app := fiber.New(configuration)
	app.Use(cors.New())
	app.Use(pprof.New())
	app.Use(middleware.ClientSecretMiddleware())

	db.Connect()
	api.SetupRoutes(app)

	log.Fatal(app.Listen(":5001"))

	defer db.Disconnect()
}
