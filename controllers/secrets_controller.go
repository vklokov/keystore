package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type SecretsController struct {
	BaseController
}

func newSecretsController() *SecretsController {
	return &SecretsController{}
}

func (self *SecretsController) All(ctx *fiber.Ctx) error {
	// user := ctx.Locals(utils.CURRENT_USER).(*entities.User)
	// secrets, err := repos.Secrets().ForUser(user.ID)

	// if err != nil {
	// 	panic(err)
	// }

	// payload := []map[string]interface{}{}
	// for _, item := range *secrets {
	// 	payload = append(payload, item.ToJson())
	// }

	// return self.responseWith200(ctx, fiber.Map{
	// 	"secrets": payload,
	// })
	return nil
}
