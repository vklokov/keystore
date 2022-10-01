package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vklokov/keystore/handlers"
)

func Register(app *fiber.App, h *handlers.AppHandler) {
	app.Get("/ping", h.HealthHandler)
}
