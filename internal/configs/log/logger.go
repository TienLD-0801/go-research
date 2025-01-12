package logger_configs

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func Logger() fiber.Handler {
	return logger.New(logger.Config{
		Format:   "[${time}] ${ip} ${method} ${path} - ${status} (${latency})\n",
		TimeZone: "Local",
	})
}
