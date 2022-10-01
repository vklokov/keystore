package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vklokov/keystore/handlers"
)

func Register(app *fiber.App, handler *handlers.AppHandler) {
	app.Get("/ping", handler.HealthHandler)
}
