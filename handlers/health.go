package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func (h *AppHandler) HealthHandler(ctx *fiber.Ctx) error {
	h.Logger.Info("Test message!!!")

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success":     true,
		"environment": os.Getenv("APP_ENV"),
	})
}
