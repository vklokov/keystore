package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"column:id"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	JTI       string `gorm:"column:jti"`
	Encrypted string `gorm:"column:encrypted"`
	Active    bool   `gorm:"column:active"`
	// Secrets   []Secret
	DeletedAt time.Time `gorm:"column:deleted_at"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (self *User) ToJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":    self.ID,
		"name":  self.Name,
		"email": self.Email,
	}

	return payload
}
