package configs

import (
	"fmt"
	"go-backend/internal/configs/database"
	"go-backend/internal/configs/exception"
	"go-backend/internal/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
)

func ServerConfig(app *fiber.App) {
	LoadEnv()

	app.Use(exception.SetCustomValidatorContext)
	database.ConnectDatabase()
	app.Use(Logger())
	routes.RegisterRoutes(app)

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	if err := app.Listen(port); err != nil {
		log.Fatal("Error starting app: ", err)
	}
}
