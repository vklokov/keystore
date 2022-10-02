package services

import (
	"errors"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/vklokov/keystore/entities"
	"github.com/vklokov/keystore/repos"
	"github.com/vklokov/keystore/utils"
)

func isUserExist(email string) bool {
	if _, err := repos.Users().FindByEmail(email); err != nil {
		return false
	}

	return true
}

func validate(params *UsersCreateParams) error {
	if len(strings.TrimSpace(params.Email)) == 0 {
		return errors.New("Email required")
	}

	if isUserExist(params.Email) {
		return errors.New("Email has already been taken")
	}

	rg := regexp.MustCompile(utils.REGEXP_EMAIL)

	if !rg.MatchString(params.Email) {
		return errors.New("Wrong email format")
	}

	if len(strings.TrimSpace(params.Name)) == 0 {
		return errors.New("Name required")
	}

	if len(strings.TrimSpace(params.Password)) == 0 {
		return errors.New("Password required")
	}

	if len(strings.TrimSpace(params.Password)) < 8 {
		return errors.New("Password should contain more than 8 charachters")
	}

	return nil
}

func UsersCreateService(params *UsersCreateParams) (*entities.User, error) {
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

	return repos.Users().Create(&user)
}
