package configs

import (
	"fmt"
	database_configs "go-backend/internal/configs/database"
	env_configs "go-backend/internal/configs/env"
	exception_configs "go-backend/internal/configs/exception"
	logger_configs "go-backend/internal/configs/log"
	"go-backend/internal/routes"
	"go-backend/internal/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
)

func ServerConfig(app *fiber.App) {
	env_configs.LoadEnv()

	database_configs.ConnectDatabase()

	app.Use(utils.AddJSONResponse,
		exception_configs.SetCustomValidatorContext,
		database_configs.AddDatabaseContext,
		logger_configs.Logger())

	routes.RegisterRoutes(app)

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	if err := app.Listen(port); err != nil {
		log.Fatal("Error starting app: ", err)
	}
}
