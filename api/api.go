package api

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {
	app.Get("/url/*", ProxYURL)
	app.Get("/m3u8/*", ProxYM3U8)
}
