package api

import (
	"github.com/MinnaSync/proxy/middleware"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/url/*", middleware.LogHeaders, ProxYURL)
	app.Get("/m3u8/*", middleware.LogHeaders, ProxYM3U8)
}
