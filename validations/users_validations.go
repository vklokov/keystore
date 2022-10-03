package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/vklokov/keystore/repos"
)

func UsersValidateUniqEmail(fl validator.FieldLevel) bool {
	if _, err := repos.Users().FindByEmail(fl.Field().String()); err != nil {
		return true
	}

	return false
}
