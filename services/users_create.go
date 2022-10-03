package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/vklokov/keystore/entities"
	"github.com/vklokov/keystore/repos"
	"github.com/vklokov/keystore/utils"
	"github.com/vklokov/keystore/validations"
)

type UsersCreateService struct {
	User   *entities.User
	Params *UsersCreateParams
}

func (self *UsersCreateService) Call() (*entities.User, *validations.VaResult) {
	if err := self.validate(); err != nil {
		return nil, err
	}

	self.persist()

	return self.User, nil
}

func (self *UsersCreateService) validate() *validations.VaResult {
	v := validator.New()
	v.RegisterValidation("uniq", validations.UsersValidateUniqEmail)

	return validations.Validate(self.Params, v)
}

func (self *UsersCreateService) persist() {
	self.User.Name = self.Params.Name
	self.User.Email = self.Params.Email
	self.User.Active = true
	self.User.Encrypted = utils.EncryptString(self.Params.Password)
	self.User.JTI = uuid.New().String()

	repos.Users().Create(self.User)
}
