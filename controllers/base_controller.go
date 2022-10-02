package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vklokov/keystore/middlewares"
)

type BaseController struct {
	AuthController    *AuthController
	PingController    *PingController
	UsersController   *UsersController
	SecretsController *SecretsController
}

func (self *BaseController) responseWith200(ctx *fiber.Ctx, payload fiber.Map, options fiber.Map) error {
	success := true

	if value, ok := options["success"]; ok {
		success = value.(bool)
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": success,
		"payload": payload,
	})
}

func (self *BaseController) responseWith400(ctx *fiber.Ctx, payload fiber.Map) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
		"success": false,
		"error":   payload,
	})
}

func (self *BaseController) responseWith401(ctx *fiber.Ctx, payload fiber.Map) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
		"success": false,
		"error":   payload,
	})
}

func (self *BaseController) responseWith422(ctx *fiber.Ctx, payload fiber.Map) error {
	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
		"success": false,
		"error":   payload,
	})
}

func Init(app *fiber.App) {
	c := BaseController{
		PingController:    newPingController(),
		UsersController:   newUsersController(),
		SecretsController: newSecretsController(),
		AuthController:    newAuthController(),
	}

	app.Get("/api/v1/ping", c.PingController.Ping)

	app.Post("/api/v1/auth", c.AuthController.SignIn)
	app.Delete("/api/v1/auth", c.AuthController.SignOut)

	app.Get("/api/v1/users", middlewares.Auth, c.UsersController.Me)
	app.Post("/api/v1/users", c.UsersController.Create)

	app.Get("/api/v1/secrets", middlewares.Auth, c.SecretsController.All)
}
