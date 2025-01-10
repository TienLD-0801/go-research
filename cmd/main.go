package main

import (
	"go-backend/internal/configs"
	"go-backend/internal/routes"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive:   true,
		StrictRouting:   false,                  // When enabled, the router treats "/foo" and "/foo/" as different.
		StructValidator: configs.NewValidator(), // Register custom struct validator
		ServerHeader:    "Backend-Research",
		AppName:         "App v0.0.1",
	})

	app.Use(logger.New(logger.Config{
		Format:   "[${time}] ${ip} ${method} ${path} - ${status} (${latency})\n",
		TimeZone: "Local",
	}))

	routes.RegisterRoutes(app)

	log.Fatal(app.Listen(":4000"))
}
