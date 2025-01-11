package users_services

import (
	"go-backend/internal/configs/database"
	"go-backend/internal/configs/exception"
	users_model "go-backend/internal/models/users"

	"github.com/gofiber/fiber/v3"
)

func CreateUser(c fiber.Ctx) error {
	user := new(users_model.UserDTO)

	cv := exception.GetCustomValidatorContext(c)

	if err := c.Bind().Body(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid data"})
	}

	errors := cv.Validate(user)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}

	newUser := users_model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User created"})
}

func DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")

	var user users_model.User

	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User deleted"})
}
