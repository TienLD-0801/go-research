package routes

import (
	global "go-backend/internal/handlers/global"
	user "go-backend/internal/handlers/users"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", global.Index)

	// User routes
	userGroup := app.Group("/user")
	userGroup.Get("/", user.UserHandler)
	userGroup.Post("/", user.CreateUser)
	userGroup.Delete("/:id", user.DeleteUser)
}
