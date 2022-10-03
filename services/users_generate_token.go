package services

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/vklokov/keystore/entities"
	"github.com/vklokov/keystore/repos"
	"github.com/vklokov/keystore/utils"
)

type UsersGenerateTokenService struct {
	User *entities.User
}

func (self *UsersGenerateTokenService) Call() string {
	self.updateJti()
	claims := self.createJWTClaims()
	token := utils.GenerateToken(claims)

	return token
}

func (self *UsersGenerateTokenService) updateJti() {
	attributes := map[string]interface{}{
		"jti": uuid.New().String(),
	}
	repos.Users().Update(self.User, attributes)
}

func (self *UsersGenerateTokenService) createJWTClaims() *jwt.MapClaims {
	tomorrow := time.Now().UTC().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"email": self.User.Email,
		"jti":   self.User.JTI,
		"exp":   tomorrow.Unix(),
	}

	return &claims
}
