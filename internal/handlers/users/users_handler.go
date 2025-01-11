package users

import (
	user "go-backend/internal/services/user"

	"github.com/gofiber/fiber/v3"
)

func UserHandler(c fiber.Ctx) error {
	return c.SendString("Handler response")
}

func CreateUser(c fiber.Ctx) error {
	return user.CreateUser(c)
}

func DeleteUser(c fiber.Ctx) error {
	return user.DeleteUser(c)
}
