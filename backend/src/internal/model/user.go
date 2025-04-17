package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID   int64   `gorm:"indexUnique;not null"`
	Username string  `gorm:"indexUnique;not null"`
	Password string  `gorm:"not null"`
	Email    *string `gorm:"indexUnique"`
	IsActive bool    `gorm:"not null;default:true"`
	Role     string  `gorm:"not null;default:'user';index"` // user, admin, superadmin
}

func (User) TableName() string {
	return "users"
}
