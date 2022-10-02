package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Mount(app *fiber.App) {
	app.Get("/api/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(map[string]interface{}{
			"success": true,
		})
	})

	app.Route("/api/auth", AuthController)
	app.Route("/api/v1/users", UsersController)
}
