package entities

import (
	"time"

	"gorm.io/gorm"
)

type Secret struct {
	gorm.Model
	ID        uint      `gorm:"column:id"`
	UserID    uint      `gorm:"column:user_id"`
	Name      string    `gorm:"column:name"`
	Login     string    `gorm:"column:login"`
	Password  string    `gorm:"column:password"`
	Email     string    `gorm:"column:email"`
	Website   string    `gorm:"column:website"`
	Note      string    `gorm:"column:note"`
	Pkey      string    `gorm:"column:pkey"`
	DeletedAt time.Time `gorm:"-"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (self *Secret) ToJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":       self.ID,
		"name":     self.Name,
		"login":    self.Login,
		"password": self.Password,
		"email":    self.Email,
		"website":  self.Website,
		"note":     self.Note,
		"pkey":     self.Pkey,
	}

	return payload
}
