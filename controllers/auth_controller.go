package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/vklokov/keystore/entities"
	"github.com/vklokov/keystore/repos"
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
	params := services.UsersSignParams{}

	json.Unmarshal(ctx.Body(), &params)

	user, err := repos.Users().FindByEmail(params.Email)

	if err != nil {
		return self.responseWith401(ctx, fiber.Map{
			"message": "Email or password incorrect",
		})
	}

	// Compare encrypted password
	a := []byte(user.Encrypted)
	b := []byte(utils.EncryptString(params.Password))

	if !utils.IsIdentical(a, b) {
		return self.responseWith401(ctx, fiber.Map{
			"message": "Email or password incorrect",
		})
	}

	token, err := services.UsersGenerateTokenService(user)

	if err != nil {
		return self.responseWith401(ctx, fiber.Map{
			"message": "Email or password incorrect",
		})
	}

	return self.responseWith200(ctx, fiber.Map{
		"accessToken": token,
	}, fiber.Map{})
}

func (self *AuthController) SignOut(ctx *fiber.Ctx) error {
	user := ctx.Locals(utils.CURRENT_USER).(*entities.User)
	_, err := services.UsersGenerateTokenService(user)

	if err != nil {
		return self.responseWith400(ctx, fiber.Map{})
	}

	return self.responseWith200(ctx, fiber.Map{}, fiber.Map{})
}
