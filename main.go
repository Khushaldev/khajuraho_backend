package main

import (
	"log"

	"khajuraho/backend/config"
	db "khajuraho/backend/database"
	"khajuraho/backend/middleware"
	"khajuraho/backend/router"

	"github.com/gofiber/fiber/v2"
)

func initApp() {
	// Load configuration
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Connect to DB
	err = db.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
}

func createApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Use(middleware.ClientSecretMiddleware())

	router.CreateRoutes(app)

	return app
}

func main() {
	initApp()

	app := createApp()

	port := config.AppConfig.AppPort
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server running on http://localhost:%s", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
