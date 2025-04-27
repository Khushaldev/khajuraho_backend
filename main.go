package main

import (
	"log"
	"os"

	"khajuraho/backend/api"
	"khajuraho/backend/config"
	db "khajuraho/backend/database"
	err "khajuraho/backend/middleware/error"

	_ "khajuraho/backend/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func loggerFunction() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile | log.Lmicroseconds)
	log.SetOutput(os.Stdout)
}

// @title Backend Service API
// @version 1.0
// @description API documentation for Khajuraho API Service
// @host      localhost:5001
// @BasePath  /api/v1

func main() {
	loggerFunction()
	config.LoadConfig()
	config.InitFirebase()

	configuration := fiber.Config{
		ErrorHandler:            err.Handler,
		EnableTrustedProxyCheck: true,
	}

	app := fiber.New(configuration)
	app.Use(cors.New())
	app.Use(pprof.New())

	db.Connect()
	api.SetupRoutes(app)

	log.Fatal(app.Listen(":5001"))

	defer db.Disconnect()
}
