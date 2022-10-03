package services

import (
	"github.com/vklokov/keystore/utils"
)

func UsersRegister(params *UsersCreateParams) (string, *utils.ValidationResult) {
	user, err := UsersCreateService(params)

	if err != nil {
		return "", err
	}

	return UsersGenerateTokenService(user), nil
}
