package controllers

import "github.com/gofiber/fiber/v2"

type PingController struct {
	BaseController
}

func newPingController() *PingController {
	return &PingController{}
}

func (self *PingController) Ping(ctx *fiber.Ctx) error {
	return self.responseWith200(ctx, fiber.Map{}, fiber.Map{})
}