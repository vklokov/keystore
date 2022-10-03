package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/vklokov/keystore/entities"
	"github.com/vklokov/keystore/services"
	"github.com/vklokov/keystore/utils"
)

type AuthController struct {
	BaseController
}

func newAuthController() *AuthController {
	return &AuthController{}
}

func (self *AuthController) SignIn(ctx *fiber.Ctx) error {
	params := &services.UsersSignParams{}

	if err := json.Unmarshal(ctx.Body(), params); err != nil {
		panic(err)
	}

	service := services.UsersSignInService{
		Params: params,
	}

	token, err := service.Call()

	if err != nil {
		return self.responseWith422(ctx, fiber.Map{
			"errors": err.ToJson(),
		})
	}

	return self.responseWith200(ctx, fiber.Map{
		"accessToken": token,
	})
}

func (self *AuthController) SignOut(ctx *fiber.Ctx) error {
	user := ctx.Locals(utils.CURRENT_USER).(*entities.User)
	service := services.UsersGenerateTokenService{
		User: user,
	}
	service.Call()

	return self.responseWith200(ctx, fiber.Map{})
}
