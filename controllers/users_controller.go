package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/vklokov/keystore/entities"
	"github.com/vklokov/keystore/services"
	"github.com/vklokov/keystore/utils"
)

type UsersController struct {
	BaseController
}

func newUsersController() *UsersController {
	return &UsersController{}
}

// POST: /api/v1/users
func (self *UsersController) Create(ctx *fiber.Ctx) error {
	params := new(services.UsersCreateParams)

	if err := json.Unmarshal(ctx.Body(), params); err != nil {
		panic(err)
	}

	token, err := services.UsersRegister(params)

	if err != nil {
		return self.responseWith422(ctx, fiber.Map{
			"errors": err.ToJson(),
		})
	}

	return self.responseWith200(ctx, fiber.Map{
		"accessToken": token,
	})
}

// GET: /api/v1/users
func (self *UsersController) Me(ctx *fiber.Ctx) error {
	user := ctx.Locals(utils.CURRENT_USER).(*entities.User)

	return self.responseWith200(ctx, fiber.Map{
		"user": user.ToJson(),
	})
}
