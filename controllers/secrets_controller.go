package controllers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vklokov/keystore/config"
	"github.com/vklokov/keystore/entities"
	"github.com/vklokov/keystore/repos"
	"github.com/vklokov/keystore/services"
	"github.com/vklokov/keystore/utils"
)

type SecretsController struct {
	BaseController
}

// GET: /api/v1/secrets
func (self *SecretsController) All(ctx *fiber.Ctx) error {
	user := ctx.Locals(utils.CURRENT_USER).(*entities.User)
	secrets, err := repos.Secrets().ForUser(user.ID)

	if err != nil {
		panic(err)
	}

	result := []Map{}

	for _, secret := range secrets {
		result = append(result, secret.ToJson())
	}

	return self.responseWith200(ctx, Map{"secrets": result})
}

// POST: /api/v1/secrets
func (self *SecretsController) Create(ctx *fiber.Ctx) error {
	user := ctx.Locals(utils.CURRENT_USER).(*entities.User)
	params := &services.SecretsParams{}

	if err := json.Unmarshal(ctx.Body(), params); err != nil {
		config.Logger.Info("JSON Encode has failed")
		return self.responseWith400(ctx, []Map{})
	}
	log.Printf("%v", params)
	service := &services.SecretsCreateService{
		Params: params,
		User:   user,
		Secret: &entities.Secret{},
	}

	secret, err := service.Call()

	if err != nil {
		return self.responseWith422(ctx, err.ToJson())
	}

	return self.responseWith200(ctx, Map{"secret": secret.ToJson()})
}

// GET: /api/v1/secrets/:id
func (self *SecretsController) Find(ctx *fiber.Ctx) error {
	user := ctx.Locals(utils.CURRENT_USER).(*entities.User)
	secretId, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return self.responseWith404(ctx, []Map{})
	}

	query := Map{"user_id": user.ID, "id": secretId}
	secret, err := repos.Secrets().FindByCondition(query)

	if err != nil {
		return self.responseWith404(ctx, []Map{})
	}

	return self.responseWith200(ctx, secret.ToJson())
}

// PUT /api/v1/secrets/:id
func (self *SecretsController) Update(ctx *fiber.Ctx) error {
	user := ctx.Locals(utils.CURRENT_USER).(*entities.User)
	secretId, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return self.responseWith404(ctx, []Map{})
	}

	query := Map{"user_id": user.ID, "id": secretId}
	secret, err := repos.Secrets().FindByCondition(query)

	if err != nil {
		return self.responseWith404(ctx, []Map{})
	}

	params := &services.SecretsParams{}

	if err := json.Unmarshal(ctx.Body(), params); err != nil {
		config.Logger.Info("JSON Encode has failed")
		return self.responseWith400(ctx, []Map{})
	}

	service := &services.SecretsUpdateService{
		User:   user,
		Params: params,
		Secret: secret,
	}

	_, vErrors := service.Call()

	if vErrors != nil {
		return self.responseWith422(ctx, vErrors.ToJson())
	}

	return self.responseWith200(ctx, secret.ToJson())
}

// DELETE /api/v1/secrets/:id
func (self *SecretsController) Delete(ctx *fiber.Ctx) error {
	return nil
}
