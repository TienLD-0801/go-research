package user_router

import (
	users_handler "go-backend/app/handlers/users"

	"github.com/gofiber/fiber/v3"
)

func UserRouter(app *fiber.App) {
	users_handler.InitUserHandler()

	userGroup := app.Group("/user")
	userGroup.Get("/", users_handler.GetUser)
	userGroup.Get("/:id", users_handler.GetUserById)
	userGroup.Post("/", users_handler.CreateUser)
	userGroup.Delete("/:id", users_handler.DeleteUser)
}
