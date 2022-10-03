package services

import (
	"github.com/vklokov/keystore/validations"
)

func UsersRegister(params *UsersCreateParams) (string, *validations.VaResult) {
	user, err := UsersCreateService(params)

	if err != nil {
		return "", err
	}

	return UsersGenerateTokenService(user), nil
}
