package db

import "time"

type User struct {
	ID        uint      `gorm:"column:id;primaryKey" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Email     string    `gorm:"column:email" json:"email"`
	Salt      string    `gorm:"column:salt"`
	Secret    string    `gorm:"column:secret"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP(6)"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP(6)"`
}
