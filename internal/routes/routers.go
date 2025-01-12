package routes

import (
	global_handler "go-backend/app/handlers/global"
	auth_router "go-backend/internal/routes/auth"
	user_router "go-backend/internal/routes/user"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", global_handler.Index)
	auth_router.AuthRouter(app)
	user_router.UserRouter(app)
}
