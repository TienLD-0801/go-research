package utils

import (
	"go-backend/internal/constants"

	"github.com/gofiber/fiber/v3"
)

func AddJSONResponse(c fiber.Ctx) error {
	c.Locals(constants.JSONResponse, JSONResponse)
	return c.Next()
}

func JSONResponse(c fiber.Ctx, status int, message string, data interface{}) error {
	return c.JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
