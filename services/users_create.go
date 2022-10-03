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
	if err := validate(params); err != nil {
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

func validate(params *UsersCreateParams) *validations.VaResult {
	v := validator.New()
	v.RegisterValidation("uniq", validations.UsersValidateUniqEmail)
	return validations.Validate(params, v)
}
