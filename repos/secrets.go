package repos

import (
	"github.com/vklokov/keystore/db"
	"github.com/vklokov/keystore/entities"
)

type SecretsRepo struct{}

func Secrets() *SecretsRepo {
	return &SecretsRepo{}
}

func (self *SecretsRepo) ForUser(userId uint) ([]*entities.Secret, error) {
	secrets := []*entities.Secret{}

	if result := db.Conn.
		Where("user_id = ?", userId).
		Find(&secrets); result.Error != nil {
		return []*entities.Secret{}, result.Error
	}

	return secrets, nil
}

func (self *SecretsRepo) FindByCondition(conditions map[string]interface{}) (*entities.Secret, error) {
	secret := &entities.Secret{}

	if result := db.Conn.
		Where(conditions).
		First(secret); result.Error != nil {
		return nil, result.Error
	}

	return secret, nil
}

func (self *SecretsRepo) Create(secret *entities.Secret) bool {
	if result := db.Conn.Create(secret); result.Error != nil {
		panic(result.Error)
	}

	return true
}

func (self *SecretsRepo) Update(secret *entities.Secret) bool {
	if result := db.Conn.Updates(secret); result.Error != nil {
		panic(result.Error)
	}

	return true
}
