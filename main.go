package main

import (
	"os"

	"github.com/MinnaSync/proxy/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "",
		AllowMethods: "GET,OPTIONS",
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	api.Register(app)

	_ = app.Listen(":" + port)
}
