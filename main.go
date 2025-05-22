package main

import (
	"strings"
	"time"

	"github.com/MinnaSync/proxy/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowMethods: "GET,OPTIONS",
		AllowOriginsFunc: func(origin string) bool {
			return strings.HasSuffix(origin, ".gura.sa.com")
		},
	}))
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("noCache") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	api.Register(app)

	_ = app.Listen(":8080")
}
