package auth_handler

import (
	auth_service "go-backend/app/services/auth"

	"github.com/gofiber/fiber/v3"
)

var authService auth_service.IAuthService

func InitAuthHandler() {
	authService = &auth_service.AuthService{}
}

func Login(c fiber.Ctx) error {
	return authService.Login(c)
}
