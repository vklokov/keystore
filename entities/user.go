package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint64 `gorm:"column:id;primaryKey;autoIncrement:true"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email;uniqueIndex"`
	JTI       string `gorm:"column:jti;index"`
	Encrypted string `gorm:"column:encrypted"`
	Active    bool   `gorm:"column:active;default:false"`
}

func (self *User) ToJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":    self.ID,
		"name":  self.Name,
		"email": self.Email,
	}

	return payload
}
