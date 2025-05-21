package model

import (
	"fmt"

	"gorm.io/gorm"
)

type ModelChannel struct {
	gorm.Model
	Name         string   `gorm:"not null"`
	Type         uint     `gorm:"default:0"`
	Status       uint8    `gorm:"default:0"`
	BaseURL      string   `gorm:"column:base_url"`
	SecretKey    string   `gorm:"not_null"`
	Tag          *string  `gorm:"indexUnique"`
	Proxy        *string  `gorm:"default:null"`
	DeleteModels []string `gorm:"type:text[]"`
	Models       []Model
}

func (ModelChannel) TableName() string {
	return "model_channels"
}

func (m *ModelChannel) BeforeDelete(tx *gorm.DB) (err error) {
	if m.ID == 0 {
		// If deletion is a batch method, for example, "db.Where(conds).Delete(&ModelChannel{})"
		return nil
	}
	var models []Model
	if err := tx.Model(m).Association("Models").Find(&models); err != nil {
		return fmt.Errorf("failed to find associated models: %w", err)
	}
	if len(models) > 0 {
		if err := tx.Model(&Model{}).Delete(&models).Error; err != nil {
			return fmt.Errorf("failed to delete associated models: %w", err)
		}
	}
	return nil
}
