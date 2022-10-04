package services

import (
	"github.com/vklokov/keystore/entities"
	"github.com/vklokov/keystore/validations"
)

type UsersRegisterService struct {
	Params *UsersCreateParams
	User   *entities.User
}

func (self *UsersRegisterService) Call() (string, *validations.Result) {
	self.User = &entities.User{}

	if err := self.createUser(); err != nil {
		return "", err
	}

	token := self.generateToken()

	return token, nil
}

func (self *UsersRegisterService) createUser() *validations.Result {
	service := UsersCreateService{
		Params: self.Params,
		User:   &entities.User{},
	}

	_, err := service.Call()

	return err
}

func (self *UsersRegisterService) generateToken() string {
	service := UsersGenerateTokenService{
		User: self.User,
	}

	return service.Call()
}

// func UsersRegister(params *UsersCreateParams) (string, *validations.VaResult) {
// 	createService := UsersCreateService{
// 		Params: params,
// 		User:   &entities.User{},
// 	}

// 	user, err := createService.Call()

// 	if err != nil {
// 		return "", err
// 	}

// 	generateTokenService := UsersGenerateTokenService{
// 		User: user,
// 	}
// 	token := generateTokenService.Call()

// 	return token, nil
// }
