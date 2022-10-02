package entities

import (
	"time"

	"gorm.io/gorm"
)

type Secret struct {
	gorm.Model
	ID        uint      `gorm:"column:id"`
	UserID    User      `gorm:"column:user_id"`
	Login     string    `gorm:"column:login"`
	Password  string    `gorm:"column:password"`
	Email     string    `gorm:"column:email"`
	Website   string    `gorm:"column:website"`
	Note      string    `gorm:"column:note"`
	PKey      string    `gorm:"column:pkey"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
