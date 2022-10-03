package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func EncryptString(str string) string {
	hasher := sha1.New()
	hasher.Write([]byte(str))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func GenerateToken(claims *jwt.MapClaims) string {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		panic(err)
	}

	return token
}

func DecodeToken(accessToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Token invalid")
}
