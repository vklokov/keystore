package services

import (
	"crypto/subtle"

	"github.com/vklokov/keystore/entities"
	"github.com/vklokov/keystore/repos"
	"github.com/vklokov/keystore/utils"
	"github.com/vklokov/keystore/validations"
)

type UsersSignInService struct {
	Params *UsersSignParams
	User   *entities.User
}

func (self *UsersSignInService) Call() (string, *validations.VaResult) {
	if err := self.findUser(); err != nil {
		return "", err
	}

	if err := self.validate(); err != nil {
		return "", err
	}

	service := UsersGenerateTokenService{
		User: self.User,
	}

	token := service.Call()

	return token, nil
}

func (self *UsersSignInService) findUser() *validations.VaResult {
	user, err := repos.Users().FindByEmail(self.Params.Email)

	if err != nil {
		return handleWrongPasswordError()
	}

	self.User = user

	return nil
}

func (self *UsersSignInService) validate() *validations.VaResult {
	a := []byte(self.User.Encrypted)
	b := []byte(utils.EncryptString(self.Params.Password))

	if res := subtle.ConstantTimeCompare(a, b); res == 0 {
		return handleWrongPasswordError()
	}

	return nil
}

func handleWrongPasswordError() *validations.VaResult {
	errorEntity := &validations.VaError{}
	errorEntity.Field = "Base"
	errorEntity.Kind = validations.KIND_VALIDATION
	errorEntity.Message = "Email or password incorrect"

	return validations.CreateResult(
		append([]*validations.VaError{}, errorEntity),
	)
}
