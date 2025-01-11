package main

import (
	configs "go-backend/internal/configs"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	configs.ServerConfig(app)
}
