package services

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/vklokov/keystore/entities"
	"github.com/vklokov/keystore/repos"
	"github.com/vklokov/keystore/utils"
)

func UsersGenerateTokenService(user *entities.User) (string, error) {
	attributes := map[string]interface{}{"jti": uuid.New().String()}
	if _, err := repos.Users().Update(user, attributes); err != nil {
		return "", err
	}

	tomorrow := time.Now().UTC().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"email": user.Email,
		"jti":   user.JTI,
		"exp":   tomorrow.Unix(),
	}

	return utils.GenerateToken(&claims)
}
