package handlers

import (
	"fmt"
	"go-backend/internal/configs"

	"github.com/gofiber/fiber/v3"
)

func UserHandler(c fiber.Ctx) error {
	return c.SendString("Handler response")
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(c fiber.Ctx) error {

	user := new(User)

	if err := c.Bind().Body(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	validator := c.App().Config().StructValidator

	if err := validator.Validate(user); err != nil {

		var errMsgs []string
		for _, e := range err.(*configs.ValidationErrors).Errors {
			errMsgs = append(errMsgs, fmt.Sprintf(
				e.Field,
				e.Message,
			))
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errMsgs,
		})
	}

	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"data":    user,
	})
}

func DeleteUser(c fiber.Ctx) error {
	return c.SendString("Delete user response")
}
