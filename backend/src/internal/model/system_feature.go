package model

import "gorm.io/gorm"

type SystemFeature struct {
	gorm.Model
	ID          int64   `gorm:"primaryKey;autoIncrement"`
	Key         string  `gorm:"indexUnique;not null"`
	Value       string  `gorm:"not null"`
	Description *string `gorm:"default:null"`
}

func (SystemFeature) TableName() string {
	return "system_features"
}
