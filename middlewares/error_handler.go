package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vklokov/keystore/config"
	"github.com/vklokov/keystore/utils"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	config.Logger.Error(err)

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return ctx.Status(code).JSON(fiber.Map{
		"success": false,
		"error": map[string]string{
			"message": utils.INTERNAL_SERVER_ERROR,
		},
	})
}
