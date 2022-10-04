package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type BaseController struct {
	Auth    *AuthController
	Ping    *PingController
	Users   *UsersController
	Secrets *SecretsController
}

type Map = map[string]interface{}

func (self *BaseController) responseWith200(ctx *fiber.Ctx, payload Map) error {
	return ctx.Status(fiber.StatusOK).JSON(Map{
		"success": true,
		"payload": payload,
	})
}

func (self *BaseController) responseWith400(ctx *fiber.Ctx, payload []Map) error {
	return self.responseWithError(ctx, fiber.StatusBadRequest, payload)
}

func (self *BaseController) responseWith401(ctx *fiber.Ctx, payload []Map) error {
	return self.responseWithError(ctx, fiber.StatusUnauthorized, payload)
}

func (self *BaseController) responseWith404(ctx *fiber.Ctx, payload []Map) error {
	return self.responseWithError(ctx, fiber.StatusNotFound, payload)
}

func (self *BaseController) responseWith422(ctx *fiber.Ctx, payload []Map) error {
	return self.responseWithError(ctx, fiber.StatusUnprocessableEntity, payload)
}

func (self *BaseController) responseWithError(ctx *fiber.Ctx, status int, payload []Map) error {
	return ctx.Status(status).JSON(Map{
		"success": false,
		"errors":  payload,
	})
}

func Create() *BaseController {
	return &BaseController{
		Ping:    &PingController{},
		Users:   &UsersController{},
		Secrets: &SecretsController{},
		Auth:    &AuthController{},
	}
}
