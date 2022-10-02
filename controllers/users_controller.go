package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/vklokov/keystore/services"
)

func UsersController(router fiber.Router) {
	router.Post("/", func(ctx *fiber.Ctx) error {
		params := services.UsersCreateParams{}

		json.Unmarshal(ctx.Body(), &params)

		token, err := services.UsersRegister(&params)

		if err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
		}

		return responseOK(ctx, fiber.Map{
			"payload": map[string]interface{}{
				"accessToken": token,
			},
		})
	})
}
