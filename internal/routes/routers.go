package routes

import (
	global_handlers "go-backend/internal/handlers/global"
	users_handlers "go-backend/internal/handlers/users"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", global_handlers.Index)

	// User routes
	userGroup := app.Group("/user")
	userGroup.Get("/", users_handlers.UserHandler)
	userGroup.Post("/", users_handlers.CreateUser)
	userGroup.Delete("/:id", users_handlers.DeleteUser)
}
