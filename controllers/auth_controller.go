package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/vklokov/keystore/entities"
	"github.com/vklokov/keystore/middlewares"
	"github.com/vklokov/keystore/repos"
	"github.com/vklokov/keystore/services"
	"github.com/vklokov/keystore/utils"
)

func responseInvalid(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": false,
		"error":   "Email or password incorrect",
	})
}

func responseOK(ctx *fiber.Ctx, payload fiber.Map) error {
	return ctx.Status(fiber.StatusOK).JSON(&payload)
}

func AuthController(router fiber.Router) {
	router.Post("/", func(ctx *fiber.Ctx) error {
		params := services.UsersSignParams{}

		json.Unmarshal(ctx.Body(), &params)

		user, err := repos.Users().FindByEmail(params.Email)

		if err != nil {
			return responseInvalid(ctx)
		}

		// Encrypted password compare
		a := []byte(user.Encrypted)
		b := []byte(utils.EncryptString(params.Password))

		if !utils.IsIdentical(a, b) {
			return responseInvalid(ctx)
		}

		// Generate access token
		token, err := services.UsersGenerateTokenService(user)

		if err != nil {
			return responseInvalid(ctx)
		}

		return responseOK(ctx, fiber.Map{
			"payload": map[string]interface{}{
				"accessToken": token,
			},
		})
	})

	router.Delete("/", middlewares.Auth, func(ctx *fiber.Ctx) error {
		user := ctx.Locals(utils.CURRENT_USER).(*entities.User)
		_, err := services.UsersGenerateTokenService(user)

		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		return nil
	})
}
