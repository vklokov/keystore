package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/vklokov/keystore/entities"
	"github.com/vklokov/keystore/repos"
	"github.com/vklokov/keystore/utils"
	"github.com/vklokov/keystore/validations"
)

func UsersCreateService(params *UsersCreateParams) (*entities.User, *validations.VaResult) {
	if err := validations.Validate(params, createValidator()); err != nil {
		return nil, err
	}

	user := entities.User{
		Name:      params.Name,
		Email:     params.Email,
		Active:    true,
		Encrypted: utils.EncryptString(params.Password),
		JTI:       uuid.New().String(),
	}

	repos.Users().Create(&user)

	return &user, nil
}

func createValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("uniq", validateUniqEmail)
	return v
}

func validateUniqEmail(fl validator.FieldLevel) bool {
	if _, err := repos.Users().FindByEmail(fl.Field().String()); err != nil {
		return true
	}

	return false
}
