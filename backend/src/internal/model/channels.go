package model

import (
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model
	Name         string   `gorm:"not null"`
	Type         uint     `gorm:"default:0"`
	Status       uint8    `gorm:"default:0"`
	BaseURL      string   `gorm:"column:base_url"`
	SecretKey    string   `gorm:"not_null"`
	Tag          *string  `gorm:"indexUnique"`
	Proxy        *string  `gorm:"default:null"`
	DeleteModels []string `gorm:"type:text[]"`
	Enable       bool     `gorm:"default:true"`
}

func (Channel) TableName() string {
	return "channels"
}
