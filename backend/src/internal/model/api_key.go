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

// func (r *APIKeyRecord) BeforeSave(tx *gorm.DB) (err error) {
// 	tx.Config.Logger.Info(tx.Statement.Context, "BeforeSave APIKeyRecord")
// 	return nil
// }

// func (r *APIKeyRecord) BeforeUpdate(tx *gorm.DB) (err error) {
// 	log.Sys().Debug("BeforeUpdate APIKeyRecord", "Dest", tx.Statement.Dest, "Name", r.Name)
// 	if tx.Statement.Changed("Name") {
// 		destName := tx.Statement.ReflectValue.FieldByName("Name").String()
// 		log.Sys().Debug("APIKeyRecord name changed", "Name", r.Name, "Dest", destName)
// 	} else {
// 		log.Sys().Debug("APIKeyRecord name not changed", "Name", r.Name)
// 	}
// 	return nil
// }
