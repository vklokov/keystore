package repos

import (
	"github.com/vklokov/keystore/db"
	"github.com/vklokov/keystore/entities"
)

type SecretsRepo struct{}

func Secrets() *SecretsRepo {
	return &SecretsRepo{}
}

func (self *SecretsRepo) ForUser(userId uint) (*[]entities.Secret, error) {
	secrets := []entities.Secret{}

	if result := db.Conn.
		Where("user_id = ?", userId).
		Find(&secrets); result.Error != nil {
		return &[]entities.Secret{}, result.Error
	}

	return &secrets, nil
}
