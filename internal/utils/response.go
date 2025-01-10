package utils

import "github.com/gofiber/fiber/v3"

func JSONResponse(c fiber.Ctx, status int, message string, data interface{}) error {
	return c.JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
