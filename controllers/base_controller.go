package controllers

import (
	"github.com/gofiber/fiber/v2"
	mw "github.com/vklokov/keystore/middlewares"
)

type BaseController struct {
	Auth    *AuthController
	Ping    *PingController
	Users   *UsersController
	Secrets *SecretsController
}

func (self *BaseController) responseWith200(ctx *fiber.Ctx, payload fiber.Map) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"payload": payload,
	})
}

func (self *BaseController) responseWith400(ctx *fiber.Ctx, payload fiber.Map) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
		"success": false,
		"payload": payload,
	})
}

func (self *BaseController) responseWith401(ctx *fiber.Ctx, payload fiber.Map) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
		"success": false,
		"payload": payload,
	})
}

func (self *BaseController) responseWith422(ctx *fiber.Ctx, payload fiber.Map) error {
	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
		"success": false,
		"payload": payload,
	})
}

func Init(app *fiber.App) {
	controller := BaseController{
		Ping:    &PingController{},
		Users:   &UsersController{},
		Secrets: &SecretsController{},
		Auth:    &AuthController{},
	}

	app.Get("/api/v1/ping", controller.Ping.Ping)

	app.Post("/api/v1/auth", controller.Auth.SignIn)
	app.Delete("/api/v1/auth", controller.Auth.SignOut)

	app.Get("/api/v1/users", mw.WithJWTAuth, controller.Users.Me)
	app.Post("/api/v1/users", controller.Users.Create)

	app.Get("/api/v1/secrets", mw.WithJWTAuth, controller.Secrets.All)
}
