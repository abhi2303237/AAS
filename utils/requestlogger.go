package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func InitRequestLogger(app *fiber.App) {
	app.Use(requestLogger())
}

func requestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Infof("[%s]: %s", c.Method(), c.Path())
		return c.Next()
	}
}
