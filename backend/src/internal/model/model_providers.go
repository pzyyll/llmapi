package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ModelProvider struct {
	gorm.Model
	Name       string  `gorm:"index;not null"`
	Tag        *string `gorm:"indexUnique"`
	Parameters datatypes.JSONType[any]
	Desc       *string `gorm:"default:null"`
	ChannelID  uint    `gorm:"index"`
	Channel    Channel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	ModelRoutingTargets []*ModelRoutingTarget `gorm:"polymorphic:Target;"`
}

func (ModelProvider) TableName() string {
	return "model_providers"
}
