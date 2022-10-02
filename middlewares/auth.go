package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vklokov/keystore/repos"
	"github.com/vklokov/keystore/utils"
)

func Auth(ctx *fiber.Ctx) error {
	chunks := strings.Split(ctx.Get("Authorization"), " ")
	accessToken := chunks[len(chunks)-1]

	if len(accessToken) == 0 {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	payload, err := utils.DecodeToken(accessToken)

	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	conditions := map[string]interface{}{"active": true, "email": payload["email"], "jti": payload["jti"]}
	user, err := repos.Users().FindByCondition(conditions)

	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	ctx.Locals(utils.CURRENT_USER, user)

	ctx.Next()

	return nil
}
