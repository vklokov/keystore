package repos

import (
	"github.com/vklokov/keystore/db"
	"github.com/vklokov/keystore/entities"
)

type UsersRepo struct{}

func Users() *UsersRepo {
	return &UsersRepo{}
}

func (self *UsersRepo) FindByEmail(email string) (*entities.User, error) {
	user := entities.User{}

	if result := db.Conn.
		Where("active = TRUE AND email = ?", email).
		First(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (self *UsersRepo) FindByCondition(conditions map[string]interface{}) (*entities.User, error) {
	user := entities.User{}

	if result := db.Conn.
		Where(conditions).
		First(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (self *UsersRepo) Create(user *entities.User) (*entities.User, error) {
	if result := db.Conn.Create(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (self *UsersRepo) Update(user *entities.User, attributes map[string]interface{}) (*entities.User, error) {
	if result := db.Conn.Model(&user).Updates(attributes); result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
