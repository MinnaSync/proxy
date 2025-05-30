package middleware

import (
	"github.com/MinnaSync/proxy/internal/logger"
	"github.com/gofiber/fiber/v2"
)

func LogHeaders(c *fiber.Ctx) error {
	jsonHeaders := make(map[string]string)

	c.Request().Header.VisitAll(func(key, value []byte) {
		jsonHeaders[string(key)] = string(value)
	})

	logger.Log.Debug("Request Headers", "headers", jsonHeaders)

	return c.Next()
}
