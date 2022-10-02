package entities

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	JTI       string    `gorm:"column:jti"`
	Encrypted string    `gorm:"column:encrypted"`
	Active    bool      `gorm:"column:active"`
	DeletedAt time.Time `gorm:"-"`
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
