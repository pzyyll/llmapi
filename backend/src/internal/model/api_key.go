package model

import (
	"time"

	"gorm.io/gorm"
)

type APIKeyRecord struct {
	gorm.Model
	UserID      int64      `gorm:"index;not null"`
	Name        string     `gorm:"not null"`
	Prefix      string     `gorm:"index;not null"`
	LookupKey   string     `gorm:"indexUnique;not null"`
	SecretHash  string     `gorm:"not null"`
	Salt        string     `gorm:"not null"`
	SecretBrief string     `gorm:"not null"`
	Scopes      int64      `gorm:"not null"`
	ExpiresAt   *time.Time `gorm:"default:null"`
	LastUsedAt  *time.Time `gorm:"default:null"`
}

func (APIKeyRecord) TableName() string {
	return "api_keys"
}
