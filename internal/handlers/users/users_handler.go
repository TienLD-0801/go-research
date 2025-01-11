package users_handlers

import (
	users_services "go-backend/internal/services/user"

	"github.com/gofiber/fiber/v3"
)

func UserHandler(c fiber.Ctx) error {
	return c.SendString("Handler response")
}

func CreateUser(c fiber.Ctx) error {
	return users_services.CreateUser(c)
}

func DeleteUser(c fiber.Ctx) error {
	return users_services.DeleteUser(c)
}
