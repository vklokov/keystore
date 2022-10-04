package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/vklokov/keystore/entities"
	"github.com/vklokov/keystore/repos"
	"github.com/vklokov/keystore/validations"
)

type SecretsCreateService struct {
	Params *SecretsParams
	User   *entities.User
	Secret *entities.Secret
}

func (self *SecretsCreateService) Call() (*entities.Secret, *validations.Result) {
	if err := self.validate(); err != nil {
		return nil, err
	}

	self.persist()

	return self.Secret, nil
}

func (self *SecretsCreateService) validate() *validations.Result {
	v := validator.New()

	return validations.Validate(self.Params, v)
}

func (self *SecretsCreateService) persist() {
	self.Secret.ID = self.Params.ID
	self.Secret.UserID = self.User.ID
	self.Secret.Name = self.Params.Name
	self.Secret.Email = self.Params.Email
	self.Secret.Login = self.Params.Login
	self.Secret.Password = self.Params.Password
	self.Secret.Website = self.Params.Website
	self.Secret.Note = self.Params.Note
	self.Secret.Pkey = self.Params.Pkey

	repos.Secrets().Create(self.Secret)
}
