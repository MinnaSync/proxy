package main

import (
	"github.com/MinnaSync/proxy/api"
	"github.com/MinnaSync/proxy/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.Conf.AllowedOrigins,
		AllowMethods: "GET,OPTIONS",
	}))

	api.Register(app)

	_ = app.Listen(":" + config.Conf.Port)
}
