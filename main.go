package main

import (
	"github.com/MinnaSync/proxy/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		ProxyHeader:             fiber.HeaderXForwardedProto,
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://minna.gura.sa.com",
		AllowMethods: "GET,OPTIONS",
	}))
	// app.Use(cache.New(cache.Config{
	// 	Next: func(c *fiber.Ctx) bool {
	// 		return c.Query("noCache") == "true"
	// 	},
	// 	Expiration:   30 * time.Minute,
	// 	CacheControl: true,
	// }))

	api.Register(app)

	_ = app.Listen(":8080")
}
