package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type SecretsController struct {
	BaseController
}

// GET: /api/v1/secrets
func (self *SecretsController) All(ctx *fiber.Ctx) error {
	return nil
}
