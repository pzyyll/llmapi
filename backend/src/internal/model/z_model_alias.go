package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ModelAlias struct {
	gorm.Model
	Alias       string `gorm:"uniqueIndex;not null"`
	Rate        int    `gorm:"default:0"`
	ExtraParams datatypes.JSONType[any]
	Models      []Model `gorm:"foreignKey:Alias;references:Alias"`
}

func (ModelAlias) TableName() string {
	return "model_alias"
}
