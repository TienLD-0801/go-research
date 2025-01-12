package users_handler

import (
	users_service "go-backend/app/services/user"

	"github.com/gofiber/fiber/v3"
)

var userService users_service.IUserService

func InitUserHandler() {
	userService = &users_service.UserService{}
}

func GetAllUser(c fiber.Ctx) error {
	return userService.GetAllUser(c)
}

func GetUserById(c fiber.Ctx) error {
	return userService.GetUserById(c)
}

func CreateUser(c fiber.Ctx) error {
	return userService.CreateUser(c)
}

func DeleteUser(c fiber.Ctx) error {
	return userService.DeleteUser(c)
}
